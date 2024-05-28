package billing

import (
	"context"
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/ports"
)

type IBillingService interface {
	GetPlanResource(ctx context.Context, req GetPlanResReq) (*billing.PlanResources, error)
	GetPricingPlans(ctx context.Context, req GetPlansReq) (*common.Pagination[GetPlansResp], error)
	SetProjectPlan(ctx context.Context, req SetPlanReq) error
	CheckPlanLimit(ctx context.Context, planId string, usage *analytics.ResourceUtil) error
}

type BillingService struct {
	repo ports.IBillingRepository
}

func NewBillingService(repo ports.IBillingRepository) *BillingService {
	return &BillingService{
		repo: repo,
	}
}
