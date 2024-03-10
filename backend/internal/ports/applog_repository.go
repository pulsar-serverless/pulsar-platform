package ports

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/log"
)

type IAppLogRepository interface {
	GetProjectLogs(ctx context.Context, projectId string, logTypes []string, searchQuery string, pageNumber int, pageSize int) (*common.Pagination[log.AppLog], error)
	DeleteProjectLogs(ctx context.Context, projectId string) error
	CreateProjectLog(ctx context.Context, log *log.AppLog) error
}
