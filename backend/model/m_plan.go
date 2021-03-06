package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/oinume/lekcije/backend/errors"
)

const DefaultMPlanID = uint8(1)

type MPlan struct {
	ID                   uint8 `gorm:"primary_key"`
	Name                 string
	InternalName         string
	StripeTestProductID  string
	Price                int16
	NotificationInterval uint8
	ShowAd               bool
	MaxFollowingTeacher  uint8
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (*MPlan) TableName() string {
	return "m_plan"
}

type MPlanService struct {
	db *gorm.DB
}

func NewMPlanService(db *gorm.DB) *MPlanService {
	return &MPlanService{db: db}
}

func (s *MPlanService) TableName() string {
	return (&MPlan{}).TableName()
}

func (s *MPlanService) FindByPK(id uint8) (*MPlan, error) {
	plan := &MPlan{}
	if err := s.db.First(plan, &MPlan{ID: id}).Error; err != nil {
		return nil, errors.NewAnnotatedError(
			errors.CodeNotFound,
			errors.WithError(err),
			errors.WithResource(errors.NewResource("m_plan", "id", id)),
		)
	}
	return plan, nil
}
