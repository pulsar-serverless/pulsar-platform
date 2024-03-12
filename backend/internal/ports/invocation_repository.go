package ports

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

type InvocationRepository interface {
	CreateInvocation(ctx context.Context, invocation *analytics.Invocation) error
}
