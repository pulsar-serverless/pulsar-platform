package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/ports"
)

type resourceService struct {
	invocationRepo ports.InvocationRepository
}

type IResourceService interface {
	CreateResourceUtil(ctx context.Context, res *domain.RuntimeResourceObj) error
	GetInvocationResourceUtil(ctx context.Context, invocationId string) (*domain.RuntimeResource, error)
}

func NewResourceService(invocationRepo ports.InvocationRepository) *resourceService {
	return &resourceService{
		invocationRepo: invocationRepo,
	}
}
