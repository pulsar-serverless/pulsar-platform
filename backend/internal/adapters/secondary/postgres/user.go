package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/domain/user"
)

func (db *Database) GetUsers(ctx context.Context, pageSize, pageNumber int, searchQuery string) (*common.Pagination[user.User], error) {
	result := &common.Pagination[user.User]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	var count int64
	db.conn.Model(&project.Project{}).Distinct("user_id").Where("user_id ILIKE ?", "%"+searchQuery+"%").Count(&count)

	var rows []*user.User
	response := db.conn.Raw(`
		SELECT user_id AS "UserId", Count(*) AS "ProjectCount" 
		FROM projects 
		WHERE user_id ILIKE ?
		GROUP BY user_id
		LIMIT ?
		OFFSET  ?
	`, "%"+searchQuery+"%", pageSize, (pageNumber-1)*pageSize).Scan(&rows)

	result.TotalPages = int64(math.Ceil(float64(count) / float64(pageSize)))
	result.Rows = rows
	return result, response.Error
}
