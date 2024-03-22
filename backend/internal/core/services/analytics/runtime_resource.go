package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"
)

func (service *resourceService) CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, proj *project.Project) error {
	resource := domain.NewResourceMetric(res, proj)

	return service.invocationRepo.CreateResourceUtil(ctx, resource)

}

func (service *resourceService) GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error) {
	return service.invocationRepo.GetInvocationResourceUtil(ctx, invocationId)
}
