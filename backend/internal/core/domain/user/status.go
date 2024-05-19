package user

import "gorm.io/gorm"

type Status = string

const (
	Active    Status = "Active"
	Suspended Status = "Suspended"
)

type AccountStatus struct {
	gorm.Model
	UserId string `gorm:"primary_key"`
	Status Status ``
	Reason string ``
}
