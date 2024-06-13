package ports

import (
	"context"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/common"
)

type IBillingRepository interface {
	GetPlanResource(ctx context.Context, planId string) (*billing.PlanResources, error)
	GetPricingPlans(ctx context.Context, pageNumber int, pageSize int) (*common.Pagination[billing.PricingPlan], error)
	SetProjectPlan(ctx context.Context, projectId string, planId string) error
	GetResourcePricing(ctx context.Context) (*billing.ResourcePricing, error)
	GetInvoice(ctx context.Context, projectId, month string) (*billing.Invoice, error)
	SaveInvoice(ctx context.Context, invoice *billing.Invoice) error
}
