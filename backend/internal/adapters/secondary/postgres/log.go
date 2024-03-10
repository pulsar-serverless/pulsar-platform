package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/log"

	"gorm.io/gorm"
)

func FilterLogs(projectId string, logTypes []string, searchQuery string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var conditions = make(map[string]interface{})
		conditions["project_id"] = projectId

		if len(logTypes) != 0 {
			conditions["type"] = logTypes
		}

		if searchQuery == "" {
			return db.Where(conditions)
		}

		return db.Where(" message ILIKE ? ", "%"+searchQuery+"%").Where(conditions)
	}
}

func (db *Database) GetProjectLogs(ctx context.Context, projectId string, logTypes []string, searchQuery string, pageNumber int, pageSize int) (*common.Pagination[log.AppLog], error) {
	var count int64
	var logs []*log.AppLog

	db.conn.Model(&log.AppLog{}).Scopes(FilterLogs(projectId, logTypes, searchQuery)).Count(&count)

	pagination := &common.Pagination[log.AppLog]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	result := db.conn.Scopes(FilterLogs(projectId, logTypes, searchQuery), Paginate(pagination)).
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

func (db *Database) DeleteProjectLogs(ctx context.Context, projectId string) error {
	result := db.conn.Where(&log.AppLog{ProjectID: projectId}).Delete(&log.AppLog{})
	return result.Error
}
