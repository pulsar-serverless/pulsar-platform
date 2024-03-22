package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/common"
)

func (db *Database) CreateResourceUtil(ctx context.Context, res *analytics.RuntimeResource) error {
	result := db.conn.Create(res)
	return result.Error
}

func (db *Database) GetProjectResourceUtil(ctx context.Context, projectId string, pageNumber int, pageSize int) (*common.Pagination[analytics.RuntimeResource], error) {
	var result []*analytics.RuntimeResource

	pagination := &common.Pagination[analytics.RuntimeResource]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	var count int64
	db.conn.Model(&analytics.RuntimeResource{}).Where(&analytics.RuntimeResource{ProjectId: projectId}).Count(&count)

	data := db.conn.Scopes(Paginate(pagination)).
		Where(&analytics.RuntimeResource{ProjectId: projectId}).
		Find(&result)

	pagination.Rows = result
	pagination.TotalPages = int64(math.Ceil(float64(count) / float64(pageSize)))

	return pagination, data.Error
}

func (db *Database) GetTotalProjectResourceUtil(ctx context.Context, projectId string) (*analytics.ResourceUtil, error) {
	var result *analytics.ResourceUtil

	data := db.conn.Raw(`
	SELECT project_id,
		SUM(total_memory) AS mem_usage,
		SUM(total_bandwidth) AS net_usage
	FROM runtime_resources
	WHERE project_id = ?
	GROUP BY project_id;
	`, projectId).Scan(&result)

	return result, data.Error
}

func (db *Database) GetMonthlyProjectResourceUtil(ctx context.Context, projectId string, month string) (*analytics.ResourceUtil, error) {
	var result *analytics.ResourceUtil

	data := db.conn.Raw(`
	WITH project_monthly_usage AS (
		SELECT project_id, total_memory, total_bandwidth,
		  CONCAT(
			EXTRACT(MONTH FROM usage_time)::text,  
			EXTRACT(YEAR FROM usage_time)::text
		  ) AS usage_month
		FROM runtime_resources
		WHERE project_id = ?
	) 
	SELECT project_id, usage_month, 
			SUM(total_memory) AS mem_usage,
			SUM(total_bandwidth) AS net_usage
	FROM project_monthly_usage
	WHERE usage_month = ?
	GROUP BY project_id, usage_month;
	`, projectId, month).Scan(&result)

	return result, data.Error
}
