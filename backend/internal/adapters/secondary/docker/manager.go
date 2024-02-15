package docker

import (
	"context"
	"io"
	"pulsar/internal/core/domain/project"

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

func (cm *ContainerManager) BuildImage(ctx context.Context, buildContext io.Reader, project *project.Project) error {
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
		return err
	}

	defer buildResponse.Body.Close()
	return checkError(buildResponse.Body)
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
