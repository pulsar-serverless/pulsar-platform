package analytics

import (
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/google/uuid"
)

type RuntimeResource struct {
	Id             uuid.UUID
	Project        *project.Project `gorm:"foreignKey:ProjectId"`
	ProjectId      string           `gorm:"index;column:project_id"`
	ContainerId    string
	TotalMemory    int64
	TotalBandwidth int64
	UsageTime      time.Time
}

func NewResourceMetric(res *RuntimeResourceObj, proj *project.Project) *RuntimeResource {
	return &RuntimeResource{
		Id:             uuid.New(),
		ContainerId:    proj.ContainerId,
		ProjectId:      proj.ID,
		TotalMemory:    res.MaxMemory,
		TotalBandwidth: res.TotalNetworkBytes,
		UsageTime:      time.Now(),
	}
}

type ResourceUtil struct {
	ProjectId   string `json:"project_id"`
	MemoryUtil  int64  `gorm:"column:mem_usage" json:"mem_usage_mb"`
	NetworkUtil int64  `gorm:"column:net_usage" json:"net_usage_mb"`
	UsagePeriod string `gorm:"column:usage_month" json:"usage_period"`
}

func ResourceUtilFromMetric(res *RuntimeResource) *ResourceUtil {
	return &ResourceUtil{
		ProjectId:   res.ProjectId,
		MemoryUtil:  res.TotalMemory, // conversion to megabyte
		NetworkUtil: res.TotalBandwidth,
		UsagePeriod: res.UsageTime.String(),
	}
}
