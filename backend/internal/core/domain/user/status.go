package user

type Status = string

const (
	Active    Status = "Active"
	Suspended Status = "Suspended"
)

type AccountStatus struct {
	UserId string `gorm:"primary_key"`
	Status Status ``
	Reason string ``
}
