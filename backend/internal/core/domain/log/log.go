package log

import (
	"time"

	"github.com/google/uuid"
)

type LogType = string

const (
	Error   LogType = "Error"
	WARNING LogType = "Warning"
	INFO    LogType = "Info"
)

type AppLog struct {
	ID        uuid.UUID `gorm:"PrimaryKey" json:"id"`
	ProjectID string    `json:"projectId"`
	Type      LogType   `gorm:"default:'Error'" json:"type"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAppLog(projectID string, logType LogType, message string) *AppLog {
	return &AppLog{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		Message:   message,
		Type:      logType,
		ProjectID: projectID,
	}
}
