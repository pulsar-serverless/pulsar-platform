package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
)

func (service *resourceService) CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj) error {
	resource := domain.NewResourceMetric(res)

	return service.invocationRepo.CreateResourceUtil(ctx, resource)

}

func (service *resourceService) GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error) {
	return service.invocationRepo.GetInvocationResourceUtil(ctx, invocationId)
}
