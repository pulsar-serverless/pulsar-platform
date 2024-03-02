package ports

import (
	"context"
	"pulsar/internal/core/domain/project"
)

type IEnvRepository interface {
	OverwriteEnvVariables(ctx context.Context, projectId string, variables []*project.EnvVariable) ([]*project.EnvVariable, error)
	GetEnvVariables(ctx context.Context, projectId string) ([]*project.EnvVariable, error)
}
