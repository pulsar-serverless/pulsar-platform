package billing

import (
	"context"
	"errors"
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/services"
)

func (billingService *BillingService) CheckPlanLimit(ctx context.Context, planId string, usage *analytics.ResourceUtil) error {
	projectPlan, err := billingService.GetPlanResource(ctx, GetPlanResReq{PlanId: planId})
	if err != nil {
		return err
	}

	if usage.MemoryUtil >= projectPlan.Memory {
		return services.NewAppError(services.ErrBadRequest, errors.New("memory usage limit reached"))
	}

	if usage.NetworkUtil >= projectPlan.Bandwidth {
		return services.NewAppError(services.ErrBadRequest, errors.New("network bandwidth limit reached"))
	}

	return nil
}
