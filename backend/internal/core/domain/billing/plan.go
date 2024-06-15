package billing

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teir string

const (
	Free       Teir = "free"
	Personal   Teir = "personal"
	Pro        Teir = "pro"
	Enterprise Teir = "enterprise"
)

type PricingPlan struct {
	ID              uuid.UUID `gorm:"PrimaryKey"`
	Name            string    `gorm:"unique"`
	Description     string    ``
	Price           float64
	PlanTeir        Teir  `gorm:"default:'free';not null;column:plan_teir"`
	NotifyThreshold int64 `gorm:"column:notify_at;default:80"`
	PlanResources   PlanResources
	gorm.Model
}
