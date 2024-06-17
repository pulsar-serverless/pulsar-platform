package billing

import (
	"time"

	"github.com/google/uuid"
)

// probably duplicated needs fixings
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
	memUtil, netUtil int64,
	requests int,
	pricing *ResourcePricing) *InvoicePriceData {
	memUsage := memUtil / MBinBytes
	memUsagePrice := float64(memUsage) * pricing.MemPrice

	netUsage := netUtil / MBinBytes
	netUsagePrice := float64(netUsage) * pricing.NetPrice

	invocationUsage := requests
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
	ID          string    `gorm:"primaryKey"`
	CreatedAt   time.Time ``
	ProjectID   string    `gorm:"not null"`
	ProjectName string    `gorm:"not null"`
	UsageMonth  string    `gorm:"not null"`
	MemUsage    int64     `gorm:"not null"`
	MemPrice    float64   `gorm:"default:0"`
	NetUsage    int64     `gorm:"not null"`
	NetPrice    float64   `gorm:"default:0"`
	Requests    int64     `gorm:"default:0"`
	ReqPrice    float64   `gorm:"default:0"`
	TotalPrice  float64   `gorm:"default:0"`
	FilePath    string    ``
}

func NewInvoice(projID, projName string, prices *InvoicePriceData, month string, totalPrice float64) *Invoice {

	return &Invoice{
		ID:          uuid.New().String(),
		CreatedAt:   time.Now(),
		ProjectID:   projID,
		ProjectName: projName,
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
