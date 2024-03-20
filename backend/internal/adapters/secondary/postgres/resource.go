package postgres

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

func (db *Database) CreateResourceUtil(ctx context.Context, res *analytics.RuntimeResource) error {
	result := db.conn.Create(res)
	return result.Error
}

func (db *Database) GetInvocationResourceUtil(ctx context.Context, containerId string) (*analytics.RuntimeResource, error) {
	return nil, nil
}
