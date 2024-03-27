package project

import (
	"pulsar/internal/core/domain/log"
	"time"

	"github.com/google/uuid"
)

type DeploymentStatus string

const (
	Ready    DeploymentStatus = "ready"
	Failed   DeploymentStatus = "failed"
	Building DeploymentStatus = "building"
	Done     DeploymentStatus = "done"
	None     DeploymentStatus = "none"
)

type Project struct {
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
	SourceCode       *SourceCode      ``
	EnvVariables     []*EnvVariable   `gorm:"foreignKey:ProjectID"`
	Logs             []*log.AppLog    `gorm:"foreignKey:ProjectID;references:ID"`
}
