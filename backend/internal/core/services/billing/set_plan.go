package billing

import (
	"context"
)

type SetPlanReq struct {
	ProjectId string `param:"projectId"`
	PlanId    string `body:"planId"`
}

func (billingService *BillingService) SetProjectPlan(ctx context.Context, req SetPlanReq) error {
	return billingService.repo.SetProjectPlan(ctx, req.ProjectId, req.PlanId)
}
