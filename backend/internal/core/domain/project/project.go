package project

import (
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
	ID               uuid.UUID        `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name             string           `gorm:"unique;not null"`
	ContainerId      string           `gorm:"unique;default:null"`
	Port             uint             `gorm:"unique;default:null"`
	ApiKey           string           ``
	DeploymentStatus DeploymentStatus `gorm:"default:'none'"`
	Subdomain        string           `gorm:"unique;default:null"`
	CreatedAt        time.Time        `gorm:"autoCreateTime"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime"`
}
