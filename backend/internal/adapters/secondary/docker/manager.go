package docker

import (
	"context"
	"errors"
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
		PublishAllPorts: true,
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

func (cm *ContainerManager) StartContainer(ctx context.Context, containerId string) (*nat.PortBinding, error) {
	err := cm.client.ContainerStart(ctx, containerId, container.StartOptions{})
	if err != nil {
		return nil, err
	}

	container, err := cm.client.ContainerInspect(ctx, containerId)
	portMapping, ok := container.NetworkSettings.Ports["3000/tcp"]

	fmt.Printf("dagem %+v\n", container.NetworkSettings.Ports)
	if !ok || len(portMapping) == 0 {
		return nil, errors.New("unable to find application port")
	}
	return &portMapping[0], err
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
	close(monitor.Stop)
	monitor.Wg.Wait()
}
