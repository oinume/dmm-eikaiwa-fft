package notifier

import (
	"fmt"
	"net/http"
	"time"

	"github.com/oinume/lekcije/server/bootstrap"
	"github.com/oinume/lekcije/server/emailer"
	"github.com/oinume/lekcije/server/errors"
	"github.com/oinume/lekcije/server/fetcher"
	"github.com/oinume/lekcije/server/logger"
	"github.com/oinume/lekcije/server/model"
	"github.com/pkg/profile"
	"go.uber.org/zap"
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
	switch *m.ProfileMode {
	case "block":
		defer profile.Start(profile.ProfilePath("."), profile.BlockProfile).Stop()
	case "cpu":
		defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()
	case "mem":
		defer profile.Start(profile.ProfilePath("."), profile.MemProfile).Stop()
	case "trace":
		defer profile.Start(profile.ProfilePath("."), profile.TraceProfile).Stop()
	}

	bootstrap.CheckCLIEnvVars()
	startedAt := time.Now().UTC()
	//if *m.LogLevel != "" {
	//	//logger.App.SetLevel(logger.NewLevel(*m.LogLevel))
	//}
	logger.App.Info("notifier started")
	defer func() {
		elapsed := time.Now().UTC().Sub(startedAt) / time.Millisecond
		logger.App.Info("notifier finished", zap.Int("elapsed", int(elapsed)))
	}()

	dbLogging := *m.LogLevel == "debug"
	db, err := model.OpenDB(bootstrap.CLIEnvVars.DBURL(), 1, dbLogging)
	if err != nil {
		return err
	}
	defer db.Close()

	if *m.NotificationInterval == 0 {
		return fmt.Errorf("-notification-interval is required")
	}
	users, err := model.NewUserService(db).FindAllEmailVerifiedIsTrue(*m.NotificationInterval)
	if err != nil {
		return err
	}
	mCountries, err := model.NewMCountryService(db).LoadAll()
	if err != nil {
		return errors.InternalWrapf(err, "Failed to load all MCountries")
	}
	fetcher := fetcher.NewLessonFetcher(nil, *m.Concurrency, *m.FetcherCache, mCountries, logger.App)

	var sender emailer.Sender
	if *m.SendEmail {
		sender = emailer.NewSendGridSender(http.DefaultClient)
	} else {
		sender = &emailer.NoSender{}
	}

	n := NewNotifier(db, fetcher, *m.DryRun, sender)
	defer n.Close()
	for _, user := range users {
		if err := n.SendNotification(user); err != nil {
			return err
		}
	}

	return nil
}
