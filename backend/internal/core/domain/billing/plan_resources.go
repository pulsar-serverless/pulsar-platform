package billing

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlanResources struct {
	ID            uuid.UUID `gorm:"PrimaryKey"`
	PricingPlanID string    ``
	Name          string    ``
	Memory        int64     `gorm:"default:102400;not null"`
	Bandwidth     int64     `gorm:"default:102400;not null"`
	Requests      int64     `gorm:"default:1000000;not null"`
	gorm.Model
}
