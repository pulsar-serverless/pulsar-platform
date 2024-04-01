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
	ID          uuid.UUID      `gorm:"PrimaryKey" json:"id"`
	Name        string         `gorm:"unique" json:"name"`
	Description string         ``
	PlanTeir    Teir           `gorm:"default:'free';not null;column:plan_teir"`
	ResourcesId *uuid.UUID     `gorm:"column:resources_id"`
	Resources   *PlanResources `gorm:"foreignKey:ResourcesId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
}
