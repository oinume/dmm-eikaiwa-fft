package daily_reporter

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/oinume/lekcije/server/model"
)

type Main struct {
	TargetDate *string
	LogLevel   *string
	DB         *gorm.DB
}

func (m *Main) Run() error {
	if *m.TargetDate == "" {
		return fmt.Errorf("-target-date is required")
	}
	date, err := time.Parse("2006-01-02", *m.TargetDate)
	if err != nil {
		return fmt.Errorf("invalid date format: %s", *m.TargetDate)
	}

	eventLogEmailService := model.NewEventLogEmailService(m.DB)
	stats, err := eventLogEmailService.FindStatsNewLessonNotifierByDate(date)
	if err != nil {
		return err
	}
	statUUs, err := eventLogEmailService.FindStatsNewLessonNotifierUUCountByDate(date)
	if err != nil {
		return err
	}

	values := make(map[string]*model.StatNewLessonNotifier, 100)
	for _, s := range stats {
		values[s.Event] = s
	}

	statsNewLessonNotifierService := model.NewStatsNewLessonNotifierService(m.DB)
	for _, s := range statUUs {
		v := values[s.Event]
		v.UUCount = s.UUCount
		if err := statsNewLessonNotifierService.CreateOrUpdate(v); err != nil {
			return err
		}
	}
	return nil
}