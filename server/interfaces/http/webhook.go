package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oinume/lekcije/server/context_data"
	"github.com/oinume/lekcije/server/event_logger"
	"github.com/oinume/lekcije/server/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SendGridEventValues struct {
	Timestamp int64  `json:"timestamp"`
	Event     string `json:"event"`
	Email     string `json:"email"`
	SGEventID string `json:"sg_event_id"`
	UserAgent string `json:"useragent"`
	URL       string `json:"url"` // Only when event=click
	// Custom args
	EmailType  string `json:"email_type"`
	UserID     string `json:"user_id"`
	TeacherIDs string `json:"teacher_ids"`
}

func (v *SendGridEventValues) GetUserID() uint32 {
	if id, err := strconv.ParseUint(v.UserID, 10, 32); err == nil {
		return uint32(id)
	}
	return 0
}

func (v *SendGridEventValues) LogToFile() {
	fields := []zapcore.Field{
		zap.Time("timestamp", time.Unix(v.Timestamp, 0)),
		zap.String("sgEventID", v.SGEventID),
		zap.String("email", v.Email),
	}

	var userID uint32
	if id, err := strconv.ParseUint(v.UserID, 10, 32); err == nil {
		userID = uint32(id)
	}
	if v.TeacherIDs != "" {
		fields = append(fields, zap.String("teacherIDs", v.TeacherIDs))
	}

	if v.Event == "open" || v.Event == "click" {
		fields = append(fields, zap.String("userAgent", v.UserAgent))
	}
	if v.Event == "click" {
		fields = append(fields, zap.String("url", v.URL))
	}

	event_logger.Log(userID, event_logger.CategoryEmail, v.Event, fields...)
}

func (v *SendGridEventValues) LogToDB(db *gorm.DB) error {
	eventLogEmail := &model.EventLogEmail{
		Datetime:   time.Unix(v.Timestamp, 0),
		Event:      v.Event,
		EmailType:  v.EmailType,
		UserID:     v.GetUserID(),
		UserAgent:  v.UserAgent,
		TeacherIDs: v.TeacherIDs,
		URL:        v.URL,
	}
	if v.EmailType == "" {
		eventLogEmail.EmailType = model.EmailTypeNewLessonNotifier
	}
	return model.NewEventLogEmailService(db).Create(eventLogEmail)
}

func (s *server) postAPISendGridEventWebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.postAPISendGridEventWebhook(w, r)
	}
}

func (s *server) postAPISendGridEventWebhook(w http.ResponseWriter, r *http.Request) {
	values := make([]SendGridEventValues, 0, 1000)
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		InternalServerError(w, err)
		return
	}
	defer r.Body.Close()
	// datetime, user_id, event(enum), event_id(varchar), text

	for _, v := range values {
		v.LogToFile()
		v.LogToDB(context_data.MustDB(r.Context()))
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}