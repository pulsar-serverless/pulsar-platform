package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/log"
)

func (db *Database) GetProjectLogs(ctx context.Context, projectId string, pageNumber int, pageSize int) (*common.Pagination[log.AppLog], error) {
	var count int64
	var logs []*log.AppLog

	db.conn.Model(&log.AppLog{}).Where(log.AppLog{ProjectID: projectId}).Count(&count)

	pagination := &common.Pagination[log.AppLog]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	result := db.conn.Scopes(Paginate(pagination)).
		Where(log.AppLog{ProjectID: projectId}).
		Order("created_at asc").
		Find(&logs)

	pagination.Rows = logs
	pagination.TotalPages = int64(math.Ceil(float64(count) / float64(pageSize)))

	return pagination, result.Error
}

func (db *Database) CreateProjectLog(ctx context.Context, log *log.AppLog) error {
	result := db.conn.Create(log)
	return result.Error
}
