package log

import (
	"context"
	"encoding/json"
	domain "pulsar/internal/core/domain/log"
)

func (ls *logService) CreateLogEvent(ctx context.Context, newLog *domain.AppLog) error {
	// zerolog.Info().Interface("output", newLog).Msg("")
	return ls.queue.Publish(ctx, ls.createLogQueue, newLog)
}

func (ls *logService) consumeCreateLogEvent(ctx context.Context, queueName string) {
	ls.queue.Consume(ctx, queueName, func(message []byte) error {
		var log domain.AppLog

		err := json.Unmarshal(message, &log)
		if err != nil {
			return err
		}

		return ls.logRepo.CreateProjectLog(ctx, &log)
	})
}
