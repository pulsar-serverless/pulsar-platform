package billing

import "context"

type SetPlanReq struct {
	ProjectId string
	PlanId    string
}

func (billingService *BillingService) SetProjectPlan(ctx context.Context, req SetPlanReq) error {
	return billingService.repo.SetProjectPlan(ctx, req.ProjectId, req.PlanId)
}
