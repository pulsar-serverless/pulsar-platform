package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
)

func (service *analyticsService) CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, inv *domain.Invocation) error {
	resource := domain.NewResourceMetric(
		inv, res.MaxMemory, res.TotalNetworkBytes,
	)

	return service.invocationRepo.CreateResourceUtil(ctx, resource)

}

func (service *analyticsService) GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error) {
	return service.invocationRepo.GetInvocationResourceUtil(ctx, invocationId)
}
