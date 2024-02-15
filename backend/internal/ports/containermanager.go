package ports

import (
	"context"
	"io"
	"pulsar/internal/core/domain/project"
)

type IContainerManager interface {
	CreateContainer(ctx context.Context, imageName string) (string, error)
	StartContainer(ctx context.Context, containerId string) error
	StopContainer(ctx context.Context, containerId string) error
	DeleteContainer(ctx context.Context, containerId string) error
	GetStatus(ctx context.Context, containerId string) (string, error)
	BuildImage(ctx context.Context, buildContext io.Reader, project *project.Project) error
}
