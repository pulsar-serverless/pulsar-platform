package postgres

import (
	"context"
	"errors"
	"fmt"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"
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
	result := db.conn.Find(&pricingPlans)

	pagination.Rows = pricingPlans

	return pagination, result.Error
}

func (db *Database) SetProjectPlan(ctx context.Context, projectId string, planId string) error {
	result := db.conn.Model(&project.Project{}).Where("id = ?", projectId).Update("plan_id", planId)
	fmt.Println(result.RowsAffected)
	return result.Error
}

func (db *Database) GetDefaultProjectPlan(ctx context.Context) (*billing.PricingPlan, error) {
	var plan billing.PricingPlan
	result := db.conn.Model(&billing.PricingPlan{}).Where("name = ?", "Free plan").First(&plan)
	return &plan, result.Error
}

func (db *Database) GetResourcePricing(ctx context.Context) (*billing.ResourcePricing, error) {
	var res []*billing.ResourcePricing

	result := db.conn.Find(&billing.ResourcePricing{}).Scan(&res)
	if result.RowsAffected <= 0 {
		return nil, errors.New("pricing not found")
	}

	return res[0], nil
}

func (db *Database) GetInvoice(ctx context.Context, projectId, month string) (*billing.Invoice, error) {
	var invoice billing.Invoice
	result := db.conn.Where(&billing.Invoice{ProjectID: projectId, UsageMonth: month}).Find(&invoice)

	if result.RowsAffected <= 0 {
		return nil, errors.New("invoice not found")
	}

	return &invoice, result.Error
}

func (db *Database) SaveInvoice(ctx context.Context, invoice *billing.Invoice) error {
	result := db.conn.Create(invoice)
	return result.Error
}
