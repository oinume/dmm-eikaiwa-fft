package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oinume/lekcije/proto-gen/go/proto/api/v1"
	"github.com/oinume/lekcije/server/errors"
)

type NotificationTimeSpan struct {
	UserID    uint32
	Number    uint8
	FromTime  string
	ToTime    string
	CreatedAt time.Time
	from      time.Time
	to        time.Time
}

func (*NotificationTimeSpan) TableName() string {
	return "notification_time_span"
}

func (s *NotificationTimeSpan) ParseTime() error {
	f, err := time.Parse("15:04:05", s.FromTime)
	if err != nil {
		return err
	}
	s.from = f

	t, err := time.Parse("15:04:05", s.ToTime)
	if err != nil {
		return err
	}
	s.to = t
	return nil
}

func (s *NotificationTimeSpan) Within(t time.Time) bool {
	if err := s.ParseTime(); err != nil {
		return false
	}
	if (t.After(s.from) || t.Equal(s.from)) && (t.Before(s.to) || t.Equal(s.to)) {
		return true
	}
	return false
}

type NotificationTimeSpanList []*NotificationTimeSpan

func (l NotificationTimeSpanList) Within(t time.Time) bool {
	target := t
	for _, timeSpan := range l {
		if err := timeSpan.ParseTime(); err != nil {
			return false
		}
		if timeSpan.Within(target) {
			return true
		}
	}
	return false
}

type NotificationTimeSpanService struct {
	db *gorm.DB
}

func NewNotificationTimeSpanService(db *gorm.DB) *NotificationTimeSpanService {
	return &NotificationTimeSpanService{db}
}

func (s *NotificationTimeSpanService) NewNotificationTimeSpansFromPB(
	userID uint32, args []*api_v1.NotificationTimeSpan,
) []*NotificationTimeSpan {
	values := make([]*NotificationTimeSpan, 0, len(args))
	for i, v := range args {
		fromTime := fmt.Sprintf("%v:%v", v.FromHour, v.FromMinute)
		toTime := fmt.Sprintf("%v:%v", v.ToHour, v.ToMinute)
		values = append(values, &NotificationTimeSpan{
			UserID:   userID,
			Number:   uint8(i + 1),
			FromTime: fromTime,
			ToTime:   toTime,
		})
	}
	return values
}

func (s *NotificationTimeSpanService) NewNotificationTimeSpansPB(args []*NotificationTimeSpan) ([]*api_v1.NotificationTimeSpan, error) {
	values := make([]*api_v1.NotificationTimeSpan, 0, len(args))
	for _, v := range args {
		fromTime, err := time.Parse("15:04:05", v.FromTime)
		if err != nil {
			return nil, errors.InternalWrapf(err, "Invalid time format: FromTime=%v", v.FromTime)
		}
		toTime, err := time.Parse("15:04:05", v.ToTime)
		if err != nil {
			return nil, errors.InternalWrapf(err, "Invalid time format: ToTime=%v", v.ToTime)
		}
		values = append(values, &api_v1.NotificationTimeSpan{
			FromHour:   int32(fromTime.Hour()),
			FromMinute: int32(fromTime.Minute()),
			ToHour:     int32(toTime.Hour()),
			ToMinute:   int32(toTime.Minute()),
		})
	}
	return values, nil
}

func (s *NotificationTimeSpanService) FindByUserID(userID uint32) ([]*NotificationTimeSpan, error) {
	sql := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = ?`, (&NotificationTimeSpan{}).TableName())
	timeSpans := make([]*NotificationTimeSpan, 0, 10)
	if err := s.db.Raw(sql, userID).Scan(&timeSpans).Error; err != nil {
		return nil, errors.InternalWrapf(err, "FindByUserID select failed: userID=%v", userID)
	}
	return timeSpans, nil
}

func (s *NotificationTimeSpanService) UpdateAll(userID uint32, timeSpans []*NotificationTimeSpan) error {
	for _, timeSpan := range timeSpans {
		if userID != timeSpan.UserID {
			return errors.InvalidArgumentf("timeSpans userID must be same")
		}
	}

	tx := s.db.Begin()
	sql := fmt.Sprintf(`DELETE FROM %s WHERE user_id = ?`, (&NotificationTimeSpan{}).TableName())
	if err := tx.Exec(sql, userID).Error; err != nil {
		return errors.InternalWrapf(err, "UpdateAll delete failed")
	}

	for _, timeSpan := range timeSpans {
		if err := tx.Create(timeSpan).Error; err != nil {
			errors.InternalWrapf(err, "UpdateAll insert failed")
		}
	}

	if err := tx.Commit().Error; err != nil {
		return errors.InternalWrapf(err, "UpdateAll commit failed")
	}

	return nil
}
