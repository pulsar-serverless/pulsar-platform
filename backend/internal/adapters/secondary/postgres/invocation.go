package postgres

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

func (db *Database) CreateInvocation(ctx context.Context, invocation *analytics.Invocation) error {
	result := db.conn.Create(invocation)
	return result.Error
}
