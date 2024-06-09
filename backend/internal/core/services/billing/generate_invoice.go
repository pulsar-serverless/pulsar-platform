package billing

import (
	"context"
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/billing"
)

type GenerateInvoiceReq struct {
	InvoiceMonth    string
	MonthResUsage   *analytics.ResourceUtil
	MonthInvocation *analytics.InvocationCount
	ResourcePrice   *billing.ResourcePricing
}

func (billingService *BillingService) GenerateInvoice(ctx context.Context, req GenerateInvoiceReq) (string, error) {
	return "", nil
}
