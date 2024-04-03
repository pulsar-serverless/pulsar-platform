package billing

import (
	"time"

	"github.com/google/uuid"
)

type Teir string

const (
	Free       Teir = "free"
	Personal   Teir = "personal"
	Pro        Teir = "pro"
	Enterprise Teir = "enterprise"
)

type PricingPlan struct {
	ID              uuid.UUID      `gorm:"PrimaryKey"`
	Name            string         `gorm:"unique"`
	Description     string         ``
	PlanTeir        Teir           `gorm:"default:'free';not null;column:plan_teir"`
	ResourcesId     *uuid.UUID     `gorm:"column:resources_id"`
	Resources       *PlanResources `gorm:"foreignKey:ResourcesId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NotifyThreshold int64          `gorm:"column:notify_at;default:80"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
}
