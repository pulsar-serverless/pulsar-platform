package ports

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

type ResourceRepository interface {
	CreateResourceUtil(ctx context.Context, resource *analytics.RuntimeResource) error
	GetInvocationResourceUtil(ctx context.Context, utilizationId string) (*analytics.RuntimeResource, error)
}
