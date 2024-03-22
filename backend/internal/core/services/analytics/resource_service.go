package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/ports"
)

type resourceService struct {
	invocationRepo ports.InvocationRepository
}

type IResourceService interface {
	CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, proj *project.Project) error
	GetProjectResourceUtil(ctx context.Context, projectId string) ([]*domain.ResourceUtil, error)
	GetTotalProjectResourceUtil(ctx context.Context, projectId string) (*domain.ResourceUtil, error)
	GetMonthlyProjectResourceUtil(ctx context.Context, projectId string, month string) (*domain.ResourceUtil, error)
}

func NewResourceService(invocationRepo ports.InvocationRepository) *resourceService {
	return &resourceService{
		invocationRepo: invocationRepo,
	}
}
