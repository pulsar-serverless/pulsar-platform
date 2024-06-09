package project

import (
	"pulsar/internal/core/domain/log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeploymentStatus string

const (
	Failed   DeploymentStatus = "failed"
	Building DeploymentStatus = "building"
	Done     DeploymentStatus = "done"
	None     DeploymentStatus = "none"
)

type Project struct {
	gorm.Model
	ID               string           `gorm:"primaryKey"`
	Name             string           `gorm:"unique;not null"`
	UserId           string           `gorm:"not null"`
	ContainerId      string           `gorm:"unique;default:null"`
	Port             uint             `gorm:"unique;default:null"`
	TokenIssuedAt    *time.Time       ``
	DeploymentStatus DeploymentStatus `gorm:"default:'none'"`
	CreatedAt        time.Time        `gorm:"autoCreateTime"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime"`
	SourceCodeId     *uuid.UUID       ``
	SourceCode       *SourceCode      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EnvVariables     []*EnvVariable   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Logs             []*log.AppLog    `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PlanId           *uuid.UUID       ``
	PricingPlan      *string          ``
}
