package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
)

func (db *Database) GetPlanResource(ctx context.Context, planId string) (*billing.PlanResources, error) {
	planRes := billing.PlanResources{PricingPlanID: planId}
	result := db.conn.First(&planRes)

	return &planRes, result.Error
}

func (db *Database) GetPricingPlans(ctx context.Context, pageNumber int, pageSize int) (*common.Pagination[billing.PricingPlan], error) {
	var pricingPlans []*billing.PricingPlan

	pagination := &common.Pagination[billing.PricingPlan]{
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}

	var planCount int64
	db.conn.Model(&billing.PricingPlan{}).Count(&planCount)

	result := db.conn.Scopes(Paginate(pagination)).Find(pricingPlans)

	pagination.Rows = pricingPlans
	pagination.TotalPages = int64(math.Ceil(float64(planCount) / float64(pageSize)))

	return pagination, result.Error
}

func (db *Database) SetProjectPlan(ctx context.Context, projectId string, planId string) error {
	result := db.conn.Where("id = ?", projectId).Update("plan_id", planId)

	return result.Error
}