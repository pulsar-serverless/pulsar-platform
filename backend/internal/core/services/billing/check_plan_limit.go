package billing

import (
	"context"
	"errors"
	"pulsar/internal/core/services"
)

func (billingService *BillingService) CheckPlanLimit(ctx context.Context, projId, planId string) error {
	resUsage, _ := billingService.resourceService.GetTotalProjectResourceUtil(ctx, projId)

	if resUsage != nil {
		projectPlan, err := billingService.GetPlanResource(ctx, GetPlanResReq{PlanId: planId})
		if err != nil {
			return err
		}

		if resUsage.MemoryUtil >= projectPlan.Memory {
			return services.NewAppError(services.ErrBadRequest, errors.New("memory usage limit reached"))
		}

		if resUsage.NetworkUtil >= projectPlan.Bandwidth {
			return services.NewAppError(services.ErrBadRequest, errors.New("network bandwidth limit reached"))
		}
	}

	return nil
}
