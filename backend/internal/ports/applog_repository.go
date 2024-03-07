package ports

import (
	"context"
	"pulsar/internal/core/domain/log"
)

type IAppLogRepository interface {
	GetProjectLogs(ctx context.Context, projectId string, pageNumber int, pageSize int) ([]*log.AppLog, error)
	CreateProjectLog(ctx context.Context, log *log.AppLog) error
}
