package ports

import (
	"context"
	"pulsar/internal/core/domain/analytics"
)

type InvocationRepository interface {
	CreateInvocation(ctx context.Context, invocation *analytics.Invocation) error
	GetInvocationsOfLast24Hours(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error)
	GetInvocationsOfLast7Days(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error)
	GetInvocationsOfLast30Days(ctx context.Context, projectId string, status analytics.InvocationStatus) ([]*analytics.InvocationCount, error)
}
