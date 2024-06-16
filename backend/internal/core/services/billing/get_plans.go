package billing

import (
	"context"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/services"
)

type GetPlansReq struct {
	PageNumber int
	PageSize   int
}

type GetPlansResp struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"desc"`
	Teir               string `json:"teir"`
	AllocatedMemory    int64  `json:"alloc_memmory"`
	AllocatedBandwidth int64  `json:"alloc_bandwidth"`
	AllocatedRequests  int64  `json:"alloc_requests"`
}

func NewGetPlanResp(plan *billing.PricingPlan, res *billing.PlanResources) *GetPlansResp {
	return &GetPlansResp{
		ID:                 plan.ID.String(),
		Name:               plan.Name,
		Description:        plan.Description,
		Teir:               string(plan.PlanTeir),
		AllocatedMemory:    res.Memory,
		AllocatedBandwidth: res.Bandwidth,
		AllocatedRequests:  res.Requests,
	}
}

func (billingService *BillingService) GetPricingPlans(ctx context.Context, req GetPlansReq) (*common.Pagination[GetPlansResp], error) {
	plans, err := billingService.repo.GetPricingPlans(ctx, req.PageNumber, req.PageSize)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	response := &common.Pagination[GetPlansResp]{
		PageSize:   plans.PageSize,
		PageNumber: plans.PageNumber,
		TotalPages: plans.TotalPages,
	}

	response.Rows = make([]*GetPlansResp, len(plans.Rows))
	for i, item := range plans.Rows {
		plan_resource, _ := billingService.GetPlanResource(ctx, GetPlanResReq{PlanId: item.ID.String()})

		response.Rows[i] = NewGetPlanResp(item, plan_resource)
	}

	return response, nil

}
