package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/ports"
)

type resourceService struct {
	resourceRepo ports.ResourceRepository
}

type IResourceService interface {
	CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj, inv *domain.Invocation) error
	GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error)
}

func NewResourceService(resourceRepo ports.ResourceRepository) *resourceService {
	return &resourceService{
		resourceRepo: resourceRepo,
	}
}
