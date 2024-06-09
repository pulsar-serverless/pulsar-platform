package billing

import (
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"

	"github.com/google/uuid"
)

// probably duplicated needs fixing
const MBinBytes int64 = 1024 * 1024

type InvoicePriceData struct {
	MemUsage int64
	MemPrice float64
	NetUsage int64
	NetPrice float64
	Requests int64
	ReqPrice float64
}

func NewInvoicePriceData(
	resUsage *analytics.ResourceUtil,
	requests *analytics.InvocationCount,
	pricing *ResourcePricing) *InvoicePriceData {
	memUsage := resUsage.MemoryUtil / MBinBytes
	memUsagePrice := float64(memUsage) * pricing.MemPrice

	netUsage := resUsage.NetworkUtil / MBinBytes
	netUsagePrice := float64(netUsage) * pricing.NetPrice

	invocationUsage := requests.Count
	invocationUsagePrice := float64(invocationUsage) * pricing.ReqPrice

	return &InvoicePriceData{
		MemUsage: memUsage,
		MemPrice: memUsagePrice,
		NetUsage: netUsage,
		NetPrice: netUsagePrice,
		Requests: int64(invocationUsage),
		ReqPrice: invocationUsagePrice,
	}
}

type Invoice struct {
	ID          string  `gorm:"primaryKey"`
	ProjectID   string  `gorm:"not null"`
	ProjectName string  `gorm:"not null"`
	UsageMonth  string  `gorm:"not null"`
	MemUsage    int64   `gorm:"not null"`
	MemPrice    float64 `gorm:"default:0"`
	NetUsage    int64   `gorm:"not null"`
	NetPrice    float64 `gorm:"default:0"`
	Requests    int64   `gorm:"default:0"`
	ReqPrice    float64 `gorm:"default:0"`
	TotalPrice  float64 `gorm:"default:0"`
}

func NewInvoice(proj *project.Project, prices *InvoicePriceData, month string, totalPrice float64) *Invoice {

	return &Invoice{
		ID:          uuid.New().String(),
		ProjectID:   proj.ID,
		ProjectName: proj.Name,
		UsageMonth:  month,
		MemUsage:    prices.MemUsage,
		NetUsage:    prices.NetUsage,
		Requests:    prices.Requests,
		MemPrice:    prices.MemPrice,
		NetPrice:    prices.NetPrice,
		ReqPrice:    prices.ReqPrice,
		TotalPrice:  totalPrice,
	}
}
