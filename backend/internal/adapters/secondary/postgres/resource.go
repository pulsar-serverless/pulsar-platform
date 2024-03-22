package postgres

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

func (db *Database) CreateResourceUtil(ctx context.Context, res *analytics.RuntimeResource) error {
	result := db.conn.Create(res)
	return result.Error
}

func (db *Database) GetProjectResourceUtil(ctx context.Context, projectId string) ([]*analytics.ResourceUtil, error) {
	return nil, nil
}

func (db *Database) GetTotalProjectResourceUtil(ctx context.Context, projectId string) (*analytics.ResourceUtil, error) {
	return nil, nil
}
func (db *Database) GetMonthlyProjectResourceUtil(ctx context.Context, projectId string, month string) (*analytics.ResourceUtil, error) {
	return nil, nil
}
