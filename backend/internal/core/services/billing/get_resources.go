package billing

import (
	"context"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/services"
)

type GetPlanResReq struct {
	PlanId string
}

func (billingService *BillingService) GetPlanResource(ctx context.Context, req GetPlanResReq) (*billing.PlanResources, error) {
	plan, err := billingService.repo.GetPlanResource(ctx, req.PlanId)

	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	return plan, nil
}
