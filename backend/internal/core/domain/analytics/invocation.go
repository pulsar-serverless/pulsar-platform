package analytics

import (
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/google/uuid"
)

type InvocationStatus = string

const (
	Success InvocationStatus = "Success"
	Error   InvocationStatus = "Error"
)

type Invocation struct {
	Id        uuid.UUID
	Project   *project.Project `gorm:"foreignKey:ProjectId"`
	Status    InvocationStatus
	ProjectId string
	StartedAt time.Time
	EndedAt   time.Time
}

func New(ProjectId string, StartedAt time.Time, EndedAt time.Time, err InvocationStatus) *Invocation {
	return &Invocation{
		Id:        uuid.New(),
		ProjectId: ProjectId,
		StartedAt: StartedAt,
		EndedAt:   EndedAt,
		Status:    err}
}
