package postgres

import (
	"context"
	"pulsar/internal/core/domain/log"
)

func (db *Database) GetProjectLogs(ctx context.Context, projectId string, pageNumber int, pageSize int) ([]*log.AppLog, error) {
	return nil, nil
}

func (db *Database) CreateProjectLog(ctx context.Context, log *log.AppLog) error {
	result := db.conn.Create(log)
	return result.Error
}
