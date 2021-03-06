package notifier

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/oinume/lekcije/backend/emailer"
	"github.com/oinume/lekcije/backend/fetcher"
	"github.com/oinume/lekcije/backend/logger"
	"github.com/oinume/lekcije/backend/model"
	"github.com/oinume/lekcije/backend/stopwatch"
)

var helper = model.NewTestHelper()
var _ = fmt.Print

type mockSenderTransport struct {
	sync.Mutex
	called      int
	requestBody string
}

func (t *mockSenderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.Lock()
	t.called++
	defer t.Unlock()
	time.Sleep(time.Millisecond * 500)
	resp := &http.Response{
		Header:     make(http.Header),
		Request:    req,
		StatusCode: http.StatusAccepted,
		Status:     "202 Accepted",
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return resp, err
	}
	t.requestBody = string(body)
	defer req.Body.Close()
	resp.Body = ioutil.NopCloser(strings.NewReader(""))
	return resp, nil
}

func TestMain(m *testing.M) {
	db := helper.DB(nil)
	defer db.Close()
	helper.TruncateAllTables(nil)
	os.Exit(m.Run())
}

func TestTeachersAndLessons_FilterBy(t *testing.T) {
	user := helper.CreateRandomUser(t)
	timeSpans := []*model.NotificationTimeSpan{
		{UserID: user.ID, Number: 1, FromTime: "15:30:00", ToTime: "16:30:00"},
		{UserID: user.ID, Number: 2, FromTime: "20:00:00", ToTime: "22:00:00"},
	}
	teacher := helper.CreateRandomTeacher(t)
	// TODO: table driven test
	lessons := []*model.Lesson{
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 15, 0, 0, 0, time.UTC)}, // excluded
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 16, 0, 0, 0, time.UTC)}, // included
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 17, 0, 0, 0, time.UTC)}, // excluded
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 21, 0, 0, 0, time.UTC)}, // included
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 23, 0, 0, 0, time.UTC)}, // excluded
	}
	tal := NewTeachersAndLessons(10)
	tal.data[teacher.ID] = &model.TeacherLessons{Teacher: teacher, Lessons: lessons}

	filtered := tal.FilterBy(model.NotificationTimeSpanList(timeSpans))
	if got, want := filtered.CountLessons(), 2; got != want {
		t.Fatalf("unexpected filtered lessons count: got=%v, want=%v", got, want)
	}

	wantTimes := []struct {
		hour, minute int
	}{
		{16, 0},
		{21, 0},
	}
	tl := filtered.data[teacher.ID]
	for i, wantTime := range wantTimes {
		if got, want := tl.Lessons[i].Datetime.Hour(), wantTime.hour; got != want {
			t.Errorf("unexpected hour: got=%v, want=%v", got, want)
		}
		if got, want := tl.Lessons[i].Datetime.Minute(), wantTime.minute; got != want {
			t.Errorf("unexpected minute: got=%v, want=%v", got, want)
		}
	}
}

func TestTeachersAndLessons_FilterByEmpty(t *testing.T) {
	//user := helper.CreateRandomUser()
	timeSpans := make([]*model.NotificationTimeSpan, 0)
	teacher := helper.CreateRandomTeacher(t)
	// TODO: table driven test
	lessons := []*model.Lesson{
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 15, 0, 0, 0, time.UTC)},
		{TeacherID: teacher.ID, Datetime: time.Date(2018, 1, 1, 16, 0, 0, 0, time.UTC)},
	}
	tal := NewTeachersAndLessons(10)
	tal.data[teacher.ID] = &model.TeacherLessons{Teacher: teacher, Lessons: lessons}

	filtered := tal.FilterBy(model.NotificationTimeSpanList(timeSpans))
	if got, want := filtered.CountLessons(), len(lessons); got != want {
		t.Fatalf("unexpected filtered lessons count: got=%v, want=%v", got, want)
	}

	wantTimes := []struct {
		hour, minute int
	}{
		{15, 0},
		{16, 0},
	}
	tl := filtered.data[teacher.ID]
	for i, wantTime := range wantTimes {
		if got, want := tl.Lessons[i].Datetime.Hour(), wantTime.hour; got != want {
			t.Errorf("unexpected hour: got=%v, want=%v", got, want)
		}
		if got, want := tl.Lessons[i].Datetime.Minute(), wantTime.minute; got != want {
			t.Errorf("unexpected minute: got=%v, want=%v", got, want)
		}
	}
}

func TestNotifier_SendNotification(t *testing.T) {
	db := helper.DB(t)
	appLogger := logger.NewAppLogger(os.Stdout, zapcore.DebugLevel)

	fetcherMockTransport, err := fetcher.NewMockTransport("../fetcher/testdata/3986.html")
	if err != nil {
		t.Fatalf("fetcher.NewMockTransport failed: err=%v", err)
	}
	fetcherHTTPClient := &http.Client{
		Transport: fetcherMockTransport,
	}

	t.Run("10_users", func(t *testing.T) {
		var users []*model.User
		const numOfUsers = 10
		for i := 0; i < numOfUsers; i++ {
			name := fmt.Sprintf("oinume+%02d", i)
			user := helper.CreateUser(t, name, name+"@gmail.com")
			teacher := helper.CreateRandomTeacher(t)
			helper.CreateFollowingTeacher(t, user.ID, teacher)
			users = append(users, user)
		}

		fetcher := fetcher.NewLessonFetcher(fetcherHTTPClient, 1, false, helper.LoadMCountries(t), appLogger)
		senderTransport := &mockSenderTransport{}
		senderHTTPClient := &http.Client{
			Transport: senderTransport,
		}
		sender := emailer.NewSendGridSender(senderHTTPClient, appLogger)
		n := NewNotifier(appLogger, db, fetcher, true, sender, stopwatch.NewSync().Start(), nil)

		ctx := context.Background()
		for _, user := range users {
			if err := n.SendNotification(ctx, user); err != nil {
				t.Fatalf("SendNotification failed: err=%v", err)
			}
		}
		// Wait all async requests are done
		n.Close(ctx, &model.StatNotifier{
			Datetime:             time.Now().UTC(),
			Interval:             10,
			Elapsed:              1000,
			UserCount:            uint32(len(users)),
			FollowedTeacherCount: uint32(len(users)),
		})

		//if got, want := senderTransport.called, numOfUsers; got <= want {
		//	t.Errorf("unexpected senderTransport.called: got=%v, want=%v", got, want)
		//}
	})

	t.Run("narrow_down_with_notification_time_span", func(t *testing.T) {
		user := helper.CreateRandomUser(t)
		teacher := helper.CreateRandomTeacher(t)
		helper.CreateFollowingTeacher(t, user.ID, teacher)

		notificationTimeSpanService := model.NewNotificationTimeSpanService(helper.DB(t))
		timeSpans := []*model.NotificationTimeSpan{
			{UserID: user.ID, Number: 1, FromTime: "02:00:00", ToTime: "03:00:00"},
			{UserID: user.ID, Number: 2, FromTime: "06:00:00", ToTime: "07:00:00"},
		}
		if err := notificationTimeSpanService.UpdateAll(user.ID, timeSpans); err != nil {
			t.Fatalf("UpdateAll failed: err=%v", err)
		}

		fetcher := fetcher.NewLessonFetcher(fetcherHTTPClient, 1, false, helper.LoadMCountries(t), nil)
		senderTransport := &mockSenderTransport{}
		senderHTTPClient := &http.Client{
			Transport: senderTransport,
		}
		sender := emailer.NewSendGridSender(senderHTTPClient, appLogger)
		n := NewNotifier(appLogger, db, fetcher, true, sender, stopwatch.NewSync().Start(), nil)
		if err := n.SendNotification(context.Background(), user); err != nil {
			t.Fatalf("SendNotification failed: err=%v", err)
		}

		n.Close(context.Background(), &model.StatNotifier{
			Datetime:             time.Now().UTC(),
			Interval:             10,
			Elapsed:              1000,
			UserCount:            1,
			FollowedTeacherCount: 1,
		}) // Wait all async requests are done before reading request body
		content := senderTransport.requestBody
		// TODO: table drive test
		if !strings.Contains(content, "02:30") {
			t.Errorf("content must contain 02:30 due to notification time span")
		}
		if !strings.Contains(content, "06:00") {
			t.Errorf("content must contain 06:00 due to notification time span")
		}
		if strings.Contains(content, "05:00") {
			t.Errorf("content must not contain 23:30 due to notification time span")
		}
		//fmt.Printf("content = %v\n", content)
	})
}

func TestNotifier_Close(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)
	db := helper.DB(t)
	appLogger := logger.NewAppLogger(os.Stdout, zapcore.DebugLevel)

	fetcherMockTransport, err := fetcher.NewMockTransport("../fetcher/testdata/3986.html")
	r.NoError(err, "fetcher.NewMockTransport failed")
	fetcherHTTPClient := &http.Client{
		Transport: fetcherMockTransport,
	}
	fetcher := fetcher.NewLessonFetcher(fetcherHTTPClient, 1, false, helper.LoadMCountries(t), appLogger)

	senderTransport := &mockSenderTransport{}
	senderHTTPClient := &http.Client{
		Transport: senderTransport,
	}
	sender := emailer.NewSendGridSender(senderHTTPClient, appLogger)

	user := helper.CreateRandomUser(t)
	teacher := helper.CreateTeacher(t, 3982, "Hena")
	helper.CreateFollowingTeacher(t, user.ID, teacher)

	n := NewNotifier(appLogger, db, fetcher, false, sender, stopwatch.NewSync().Start(), nil)
	err = n.SendNotification(context.Background(), user)
	r.NoError(err, "SendNotification failed")
	n.Close(context.Background(), &model.StatNotifier{
		Datetime:             time.Now().UTC(),
		Interval:             10,
		Elapsed:              1000,
		UserCount:            1,
		FollowedTeacherCount: 1,
	})

	teacherService := model.NewTeacherService(db)
	updatedTeacher, err := teacherService.FindByPK(teacher.ID)
	r.NoError(err)
	a.NotEqual(teacher.CountryID, updatedTeacher.CountryID)
	a.NotEqual(teacher.FavoriteCount, updatedTeacher.FavoriteCount)
	a.NotEqual(teacher.Rating, updatedTeacher.Rating)
	a.NotEqual(teacher.ReviewCount, updatedTeacher.ReviewCount)
}
