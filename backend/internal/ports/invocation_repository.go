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
	CreateResourceUtil(ctx context.Context, resource *analytics.RuntimeResource) error
	GetProjectResourceUtil(ctx context.Context, projectId string) ([]*analytics.ResourceUtil, error)
	GetTotalProjectResourceUtil(ctx context.Context, projectId string) (*analytics.ResourceUtil, error)
	GetMonthlyProjectResourceUtil(ctx context.Context, projectId string, month string) (*analytics.ResourceUtil, error)
}
