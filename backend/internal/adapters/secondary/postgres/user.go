package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/domain/user"

	"gorm.io/gorm/clause"
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
		SELECT 
			projects.user_id AS "UserId", 
			Count(CASE WHEN projects.deleted_at IS NULL THEN 1 END) AS "ProjectCount", 
			COALESCE(account_statuses.status, 'Active') AS "Status"
		FROM 
			projects 
		LEFT JOIN 
			account_statuses
		ON
			account_statuses.user_id =  projects.user_id
		WHERE projects.user_id ILIKE ?
		GROUP BY projects.user_id, account_statuses.status
		LIMIT ?
		OFFSET  ?
	`, "%"+searchQuery+"%", pageSize, (pageNumber-1)*pageSize).Scan(&rows)

	result.TotalPages = int64(math.Ceil(float64(count) / float64(pageSize)))
	result.Rows = rows
	return result, response.Error
}

func (db *Database) ChangeAccountStatus(ctx context.Context, userId, status string) error {
	result := db.conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status"}),
	}).Create(&user.AccountStatus{UserId: userId, Status: status})
	return result.Error
}

func (db *Database) GetAccountStatus(ctx context.Context, userId string) (string, error) {
	var status string
	result := db.conn.Model(&user.AccountStatus{}).Select("status").Where(&user.AccountStatus{UserId: userId}).Find(&status)
	return status, result.Error
}
