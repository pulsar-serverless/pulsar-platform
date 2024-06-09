package billing

import (
	"context"
	"math"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/project"
)

type GenerateInvoiceReq struct {
	Project        *project.Project
	InvoiceMonth   string
	InvoicePricing *billing.InvoicePriceData
}

func calculateTotalPrice(req GenerateInvoiceReq) float64 {
	totalPrice := req.InvoicePricing.MemPrice + req.InvoicePricing.NetPrice + req.InvoicePricing.ReqPrice

	return math.Round(totalPrice*100) / 100
}

func (billingService *BillingService) GenerateInvoice(ctx context.Context, req GenerateInvoiceReq) (string, error) {
	totalPrice := calculateTotalPrice(req)

	invoice := billing.NewInvoice(
		req.Project,
		req.InvoicePricing,
		req.InvoiceMonth,
		totalPrice,
	)

	err := billingService.repo.SaveInvoice(ctx, invoice)
	if err != nil {
		return "", nil
	}

	return invoice.ID, nil
}
