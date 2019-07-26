package notifier

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/oinume/lekcije/server/config"
	"github.com/oinume/lekcije/server/emailer"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/fetcher"
	"github.com/oinume/lekcije/server/logger"
	"github.com/oinume/lekcije/server/model"
	"github.com/oinume/lekcije/server/open_census"
	"github.com/oinume/lekcije/server/stopwatch"
	"github.com/oinume/lekcije/server/util"
	"github.com/pkg/profile"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type Main struct {
	Concurrency          *int
	DryRun               *bool
	NotificationInterval *int
	SendEmail            *bool
	FetcherCache         *bool
	LogLevel             *string
	ProfileMode          *string
}

func (m *Main) Run() error {
	sw := stopwatch.NewSync()
	sw.Start()
	var storageClient *storage.Client
	switch *m.ProfileMode {
	case "block":
		defer profile.Start(profile.ProfilePath("."), profile.BlockProfile).Stop()
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
	case "trace":
		defer profile.Start(profile.ProfilePath("."), profile.TraceProfile).Stop()
	case "stopwatch":
		var err error
		storageClient, err = newStorageClient()
		if err != nil {
			return err
		}
	}

	config.MustProcessDefault()
	startedAt := time.Now().UTC()
	//if *m.LogLevel != "" {
	//	//logger.App.SetLevel(logger.NewLevel(*m.LogLevel))
	//}
	logger.App.Info(fmt.Sprintf("notifier started (interval=%d)", *m.NotificationInterval))
	defer func() {
		elapsed := time.Now().UTC().Sub(startedAt) / time.Millisecond
		logger.App.Info(
			fmt.Sprintf("notifier finished (interval=%d)", *m.NotificationInterval),
			zap.Int("elapsed", int(elapsed)),
		)
	}()

	const serviceName = "notifier"
	exporter, flush, err := open_census.NewExporter(config.DefaultVars, serviceName)
	if err != nil {
		log.Fatalf("NewExporter failed: %v", err)
	}
	defer flush()
	trace.RegisterExporter(exporter)

	ctx, span := trace.StartSpan(context.Background(), "main")
	defer span.End()
	db, err := model.OpenDB(config.DefaultVars.DBURL(), 1, config.DefaultVars.DebugSQL)
	if err != nil {
		return err
	}
	defer db.Close()
	sw.Mark("model.OpenDB")

	//for i := 0; i < 5; i++ {
	//	doWork(ctx)
	//}

	if *m.NotificationInterval == 0 {
		return fmt.Errorf("-notification-interval is required")
	}
	users, err := model.NewUserService(db).FindAllEmailVerifiedIsTrue(ctx, *m.NotificationInterval)
	if err != nil {
		return err
	}
	mCountries, err := model.NewMCountryService(db).LoadAll(ctx)
	if err != nil {
		return err
	}
	lessonFetcher := fetcher.NewLessonFetcher(nil, *m.Concurrency, *m.FetcherCache, mCountries, logger.App)

	var sender emailer.Sender
	if *m.SendEmail {
		sender = emailer.NewSendGridSender(http.DefaultClient)
	} else {
		sender = &emailer.NoSender{}
	}

	statNotifier := &model.StatNotifier{
		Datetime:             startedAt,
		Interval:             uint8(*m.NotificationInterval),
		Elapsed:              0,
		UserCount:            uint32(len(users)),
		FollowedTeacherCount: 0,
	}
	n := NewNotifier(db, lessonFetcher, *m.DryRun, sender, sw, storageClient)
	defer n.Close(statNotifier)
	for _, user := range users {
		if err := n.SendNotification(ctx, user); err != nil {
			return err
		}
	}

	return nil
}

func newStorageClient() (*storage.Client, error) {
	serviceAccountKey := os.Getenv("GCP_SERVICE_ACCOUNT_KEY")
	if serviceAccountKey == "" {
		return nil, errors.NewInternalError(errors.WithMessage("Env not found"))
	}
	f, err := util.GenerateTempFileFromBase64String("", "gcp-", serviceAccountKey)
	if err != nil {
		return nil, err
	}
	defer func() {
		os.Remove(f.Name())
	}()
	return storage.NewClient(context.Background(), option.WithCredentialsFile(f.Name()))
}

func doWork(ctx context.Context) {
	// 4. Start a child span. This will be a child span because we've passed
	// the parent span's ctx.
	_, span := trace.StartSpan(ctx, "doWork")
	// 5a. Make the span close at the end of this function.
	defer span.End()

	fmt.Println("doing busy work")
	time.Sleep(80 * time.Millisecond)
	buf := bytes.NewBuffer([]byte{0xFF, 0x00, 0x00, 0x00})
	num, err := binary.ReadVarint(buf)
	if err != nil {
		// 6. Set status upon error
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeUnknown,
			Message: err.Error(),
		})
	}

	// 7. Annotate our span to capture metadata about our operation
	span.Annotate([]trace.Attribute{
		trace.Int64Attribute("bytes to int", num),
	}, "Invoking doWork")
	time.Sleep(20 * time.Millisecond)
}
