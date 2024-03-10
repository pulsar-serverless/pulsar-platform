package log

import (
	"context"
	"pulsar/internal/core/domain/common"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/ports"
)

type logService struct {
	queue          ports.IMessageQueue
	createLogQueue string
	logRepo        ports.IAppLogRepository
}

type ILogService interface {
	CreateLogEvent(ctx context.Context, newLog *domain.AppLog) error
	GetProjectLogs(ctx context.Context, request GetLogsRequest) (*common.Pagination[domain.AppLog], error)
}

func NewLogService(mq ports.IMessageQueue, logRepo ports.IAppLogRepository) *logService {
	queue := mq.CreateQueue("CREATE_LOG_QUEUE")

	service := &logService{
		queue:          mq,
		logRepo:        logRepo,
		createLogQueue: queue.Name,
	}

	go service.consumeCreateLogEvent(context.Background(), queue.Name)

	return service
}
