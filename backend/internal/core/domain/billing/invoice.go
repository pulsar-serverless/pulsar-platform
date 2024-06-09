package billing

import (
	"pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"

	"github.com/google/uuid"
)

// probably duplicated needs fixing
const MBinBytes int64 = 1024 * 1024

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
	Discount    int64   `gorm:"default:0"`
}

func NewInvoiceData(proj *project.Project, month string, resUsage *analytics.ResourceUtil, requests *analytics.InvocationCount, price *ResourcePricing, discount int64) *Invoice {
	memUsage := resUsage.MemoryUtil / MBinBytes
	memUsageTotal := float64(memUsage) * price.MemPrice

	netUsage := resUsage.NetworkUtil / MBinBytes
	netUsageTotal := float64(netUsage) * price.NetPrice

	requestTotal := float64(requests.Count) * price.ReqPrice

	return &Invoice{
		ID:          uuid.New().String(),
		ProjectID:   proj.ID,
		ProjectName: proj.Name,
		UsageMonth:  month,
		MemUsage:    memUsage,
		NetUsage:    netUsage,
		Requests:    int64(requests.Count),
		MemPrice:    memUsageTotal,
		NetPrice:    netUsageTotal,
		ReqPrice:    requestTotal,
	}
}
