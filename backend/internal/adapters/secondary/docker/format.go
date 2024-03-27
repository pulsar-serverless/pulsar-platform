package docker

import (
	"bytes"
	"encoding/json"
	resource "pulsar/internal/core/domain/analytics"

	"github.com/docker/docker/api/types"
)

func formatStats(res *resource.RuntimeResourceObj, stats types.ContainerStats, dockerStats *resource.DockerStats) (*resource.RuntimeResourceObj, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stats.Body)

	err := json.Unmarshal(buf.Bytes(), &dockerStats)
	if err != nil {
		return res, err
	}

	if dockerStats.MemoryStats.Total > res.MaxMemory {
		res.MaxMemory = dockerStats.MemoryStats.Total
	}

	return res, nil

}
