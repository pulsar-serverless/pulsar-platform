package analytics

import "github.com/google/uuid"

type RuntimeResource struct {
	Id             uuid.UUID
	Invocation     *Invocation `gorm:"foreignKey:InvocationId"`
	InvocationId   uuid.UUID
	TotalTime      int64
	TotalMemory    int64
	TotalBandwidth int64
}

func NewResourceMetric(invocation *Invocation, totalMem, totalBandwidth int64) *RuntimeResource {
	return &RuntimeResource{
		Id:             uuid.New(),
		InvocationId:   invocation.Id,
		TotalTime:      int64(invocation.EndedAt.Sub(invocation.StartedAt).Seconds()),
		TotalMemory:    totalMem,
		TotalBandwidth: totalBandwidth,
	}
}
