package billing

import (
	"time"
)

type Invoice struct {
	ID             string    `gorm:"primaryKey"`
	ProjectName    string    `gorm:"not null"`
	ProjectCreated time.Time ``
	UsageMonth     string    `gorm:"not null"`
	MemUsage       int64     `gorm:"not null"`
	MemPrice       float64   `gorm:"default:0"`
	NetUsage       int64     `gorm:"not null"`
	NetPrice       float64   `gorm:"default:0"`
	TotalPrice     float64   `gorm:"default:0"`
	Discount       int64     `gorm:"default:0"`
}
