package billing

import (
	"context"
	"errors"
	"fmt"
	"math"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"pulsar/internal/core/services/analytics"
	projService "pulsar/internal/core/services/project"
	"time"

	"github.com/go-pdf/fpdf"
)

// HTTP Request
type GenerateInvoiceReq struct {
	ProjectID string `param:"projectId"`
	Month     string `query:"month"`
}

// HTTP Response
type GenerateInvoiceResp struct {
	FilePath string `json:"file"`
}

// Intermediate DTO
type GenerateInvoiceData struct {
	Project        *project.Project
	InvoiceMonth   string
	InvoicePricing *billing.InvoicePriceData
}

func (billingService *BillingService) checkInvoiceExists(ctx context.Context, req GenerateInvoiceReq) (*billing.Invoice, error) {
	return billingService.repo.GetInvoice(ctx, req.ProjectID, req.Month)
}

func (billingService *BillingService) checkMonthBeforeCurrent(usageMonth string) bool {
	month, year := GetMonthYear(usageMonth)

	currentDate := time.Now()

	if year > currentDate.Year() {
		return false
	}

	if month >= int(currentDate.Month()) {
		return false
	}

	return true
}

func (billingService *BillingService) GenerateInvoice(ctx context.Context, req GenerateInvoiceReq) (*GenerateInvoiceResp, error) {
	canGenerate := billingService.checkMonthBeforeCurrent(req.Month)

	if !canGenerate {
		return nil, services.NewAppError(services.ErrBadRequest, errors.New("can't generate invoice for given period"))
	}

	exists, _ := billingService.checkInvoiceExists(ctx, req)
	if exists != nil {
		return &GenerateInvoiceResp{FilePath: exists.FilePath}, nil
	}

	invoiceData, err := billingService.generateInvoiceData(ctx, req.ProjectID, req.Month)
	if err != nil {
		return nil, err
	}

	proj, err := billingService.projectService.GetProject(ctx, projService.GetProjectReq{ProjectId: req.ProjectID})
	if err != nil {
		return nil, err
	}

	data := GenerateInvoiceData{
		Project:        proj,
		InvoiceMonth:   req.Month,
		InvoicePricing: invoiceData,
	}

	totalPrice := calculateTotalPrice(data)

	invoice := billing.NewInvoice(
		data.Project.ID,
		data.Project.Name,
		data.InvoicePricing,
		data.InvoiceMonth,
		totalPrice,
	)

	file, err := billingService.generatePdf(invoice)
	if err != nil {
		return nil, err
	}

	invoice.FilePath = file

	err = billingService.repo.SaveInvoice(ctx, invoice)
	if err != nil {
		return nil, nil
	}

	return &GenerateInvoiceResp{FilePath: file}, nil
}

func calculateTotalPrice(req GenerateInvoiceData) float64 {
	totalPrice := req.InvoicePricing.MemPrice + req.InvoicePricing.NetPrice + req.InvoicePricing.ReqPrice

	return math.Round(totalPrice*100) / 100
}

func (billingService *BillingService) generateInvoiceData(ctx context.Context, projId, month string) (*billing.InvoicePriceData, error) {
	resourceUsage, err := billingService.resourceService.GetMonthlyProjectResourceUtil(ctx, projId, month)
	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, errors.New("insufficient ressource data to generate invoice"))
	}

	requestUsage, err := billingService.analyticsService.GetMonthlyInvocations(ctx, analytics.GetInvocations{ProjectId: projId, Status: "Success"})
	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, errors.New("insufficient ressource data to generate invoice"))
	}

	var requestCount int
	for _, request := range requestUsage {
		requestCount += request.Count
	}

	resourcePricing, err := billingService.repo.GetResourcePricing(ctx)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, errors.New("pricing not found"))

	}

	invoiceData := billing.NewInvoicePriceData(resourceUsage.MemoryUtil, resourceUsage.NetworkUtil, requestCount, resourcePricing)

	return invoiceData, nil
}

func (billingService *BillingService) generatePdf(invoice *billing.Invoice) (string, error) {
	pdf := fpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 16)

	pdf.CellFormat(190, 10, "Invoice", "0", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "B", 24)

	pdf.CellFormat(190, 40, "Pulsar Serverless Platform", "0", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 12)

	pdf.CellFormat(100, 10, fmt.Sprintf("Invoice ID: %s", invoice.ID), "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("Date: %s", invoice.CreatedAt.Format("2006-01-02")), "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("Project ID: %s", invoice.ProjectID), "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("Project Name: %s", invoice.ProjectName), "0", 1, "L", false, 0, "")
	pdf.CellFormat(100, 10, fmt.Sprintf("Invoice Month: %s", invoice.UsageMonth), "0", 1, "L", false, 0, "")

	pdf.SetFillColor(220, 220, 220)
	pdf.CellFormat(40, 10, "Item", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "C", true, 0, "")
	pdf.CellFormat(40, 10, "Price", "1", 1, "C", true, 0, "")

	pdf.CellFormat(40, 10, "Memory Usage(MB)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("%d", invoice.MemUsage), "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", invoice.MemPrice), "1", 1, "C", false, 0, "")

	pdf.CellFormat(40, 10, "Network Usage(MB)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("%d", invoice.NetUsage), "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", invoice.NetPrice), "1", 1, "C", false, 0, "")

	pdf.CellFormat(40, 10, "Requests(Count)", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("%d", invoice.Requests), "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", invoice.ReqPrice), "1", 1, "C", false, 0, "")

	pdf.CellFormat(120, 10, "Total Price", "1", 0, "R", true, 0, "")
	pdf.CellFormat(40, 10, fmt.Sprintf("$%.2f", invoice.TotalPrice), "1", 1, "C", false, 0, "")

	return billingService.fileRepo.SaveInvoicePDF(invoice, pdf)
}
