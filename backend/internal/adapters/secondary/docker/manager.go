package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	resource "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type ContainerManager struct {
	client *client.Client
}

func NewContainerManager() *ContainerManager {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		panic("Unable to connect to docker")
	}
	return &ContainerManager{client}
}

func (cm *ContainerManager) Close() error {
	return cm.client.Close()
}

func (cm *ContainerManager) BuildImage(ctx context.Context, buildContext io.Reader, project *project.Project) (io.ReadCloser, error) {
	buildOptions := types.ImageBuildOptions{
		Tags:           []string{project.Name},
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Dockerfile:     "dockerfile",
	}

	buildResponse, err := cm.client.ImageBuild(ctx, buildContext, buildOptions)
	if err != nil {
		return nil, err
	}

	return buildResponse.Body, nil
}

func (cm *ContainerManager) CreateContainer(ctx context.Context, imageName string) (string, error) {
	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"3000/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "3000",
				},
			},
		},
	}

	containerConfig := &container.Config{
		Image: imageName,
		ExposedPorts: nat.PortSet{
			"3000/tcp": struct{}{},
		},
		Tty:          true,
		AttachStdout: true,
		AttachStderr: true,
	}

	resp, err := cm.client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")

	return resp.ID, err
}

func (cm *ContainerManager) StartContainer(ctx context.Context, containerId string) error {
	return cm.client.ContainerStart(ctx, containerId, container.StartOptions{})
}

func (cm *ContainerManager) StopContainer(ctx context.Context, containerId string) error {
	return cm.client.ContainerStop(ctx, containerId, container.StopOptions{})
}

func (cm *ContainerManager) DeleteContainer(ctx context.Context, containerId string) error {
	return cm.client.ContainerRemove(ctx, containerId, container.RemoveOptions{})
}

func (cm *ContainerManager) GetStatus(ctx context.Context, containerId string) (string, error) {
	container, err := cm.client.ContainerInspect(ctx, containerId)
	return container.State.Status, err
}

func (cm *ContainerManager) GetContainerLogs(ctx context.Context, containerId string) (io.ReadCloser, error) {
	return cm.client.ContainerLogs(ctx, containerId, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
		Timestamps: true,
		Details:    true,
	})
}

func (cm *ContainerManager) GetContainerStats(ctx context.Context, containerId string, res *resource.RuntimeResourceObj, monitor *resource.RuntimeResMonitor) error {
	monitor.Wg.Add(1)
	defer monitor.Wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var dockerStats resource.DockerStats

	// offset containerization
	time.Sleep(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			stats, err := cm.client.ContainerStats(ctx, containerId, false)
			if err != nil {
				return err
			}

			res, err = formatStats(res, stats, &dockerStats)
			if err != nil {
				return err
			}

		case <-monitor.Stop:
			stats, err := cm.client.ContainerStats(ctx, containerId, false)
			if err != nil {
				return err
			}
			res, err = formatStats(res, stats, &dockerStats)
			if err != nil {
				return err
			}

			res.TotalNetworkBytes = dockerStats.PortInterface.Recieved + dockerStats.PortInterface.Transmitted
			return nil
		}
	}
}

func (cm *ContainerManager) StopContainerStats(ctx context.Context, monitor *resource.RuntimeResMonitor) {
	fmt.Println("Inside stop container stats")
	close(monitor.Stop)
	monitor.Wg.Wait()
}

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
