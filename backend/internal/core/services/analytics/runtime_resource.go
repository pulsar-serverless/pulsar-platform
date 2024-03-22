package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

func (service *resourceService) CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, proj *project.Project) error {
	resource := domain.NewResourceMetric(res, proj)

	return service.invocationRepo.CreateResourceUtil(ctx, resource)

}

func (service *resourceService) GetProjectResourceUtil(ctx context.Context, projectId string) ([]*domain.ResourceUtil, error) {
	result, err := service.invocationRepo.GetProjectResourceUtil(ctx, projectId)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return result, nil
}

func (service *resourceService) GetTotalProjectResourceUtil(ctx context.Context, projectId string) (*domain.ResourceUtil, error) {
	result, err := service.invocationRepo.GetTotalProjectResourceUtil(ctx, projectId)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return result, nil
}

func (service *resourceService) GetMonthlyProjectResourceUtil(ctx context.Context, projectId string, month string) (*domain.ResourceUtil, error) {
	result, err := service.invocationRepo.GetMonthlyProjectResourceUtil(ctx, projectId, month)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return result, nil
}
