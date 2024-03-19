package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
)

func (service *resourceService) CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, inv *domain.Invocation) error {
	resource := domain.NewResourceMetric(
		inv, res.MaxMemory, res.TotalNetworkBytes,
	)

	return service.resourceRepo.CreateResourceUtil(ctx, resource)

}

func (service *resourceService) GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error) {
	return service.resourceRepo.GetInvocationResourceUtil(ctx, invocationId)
}
