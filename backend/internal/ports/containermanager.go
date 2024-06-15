package ports

import (
	"context"
	"io"
	resource "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/domain/project"

	"github.com/docker/go-connections/nat"
)

type IContainerManager interface {
	CreateContainer(ctx context.Context, imageName string) (string, error)
	StartContainer(ctx context.Context, containerId string) (*nat.PortBinding, error)
	StopContainer(ctx context.Context, containerId string) error
	DeleteContainer(ctx context.Context, containerId string) error
	GetStatus(ctx context.Context, containerId string) (string, error)
	BuildImage(ctx context.Context, buildContext io.Reader, project *project.Project) (io.ReadCloser, error)
	GetContainerLogs(ctx context.Context, containerId string) (io.ReadCloser, error)
	GetContainerStats(ctx context.Context, containerId string, res *resource.RuntimeResourceObj, monitor *resource.RuntimeResMonitor) error
	StopContainerStats(ctx context.Context, monitor *resource.RuntimeResMonitor)
}
