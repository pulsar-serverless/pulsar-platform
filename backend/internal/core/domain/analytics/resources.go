package analytics

import (
	"sync"

	"github.com/google/uuid"
)

type RuntimeResource struct {
	Id             uuid.UUID
	ContainerId    string
	TotalMemory    int64
	TotalBandwidth int64
}

func NewResourceMetric(res *RuntimeResourceObj) *RuntimeResource {
	return &RuntimeResource{
		Id:             uuid.New(),
		ContainerId:    res.ContainerId,
		TotalMemory:    res.MaxMemory,
		TotalBandwidth: res.TotalNetworkBytes,
	}
}

// docker resource stats
type MemoryStats struct {
	Total int64       `json:"usage"`
	Stats interface{} `json:"stats"`
}

type NetworkStatsInteface struct {
	Recieved    int64 `json:"rx_bytes"`
	Transmitted int64 `json:"tx_bytes"`
}

type NetworkStats struct {
	PortInterface NetworkStatsInteface `json:"eth0"`
}

type DockerStats struct {
	MemoryStats  MemoryStats `json:"memory_stats"`
	NetworkStats `json:"networks"`
}

type RuntimeResourceObj struct {
	ContainerId       string
	MaxMemory         int64
	TotalNetworkBytes int64
}

func NewRuntimeResObj() *RuntimeResourceObj {
	return &RuntimeResourceObj{}
}

type RuntimeResMonitor struct {
	Wg   *sync.WaitGroup
	Stop chan struct{}
}

func NewRuntimeResMonitor() *RuntimeResMonitor {
	return &RuntimeResMonitor{
		Wg:   new(sync.WaitGroup),
		Stop: make(chan struct{}),
	}
}
