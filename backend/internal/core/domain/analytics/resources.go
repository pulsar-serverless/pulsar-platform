package analytics

import (
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/google/uuid"
)

type RuntimeResource struct {
	Id             uuid.UUID
	Project        *project.Project `gorm:"foreignKey:ProjectId"`
	ProjectId      string
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
	ProjectId   string    `json:"project_id"`
	MemoryUtil  int64     `json:"mem_usage"`
	NetworkUtil int64     `json:"net_usage"`
	UsageTime   time.Time `json:"usage_time"`
}
