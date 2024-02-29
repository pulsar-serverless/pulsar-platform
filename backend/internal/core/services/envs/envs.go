package envs

import (
	"context"
	"pulsar/internal/core/domain/project"
	services "pulsar/internal/core/services/project"
	"pulsar/internal/ports"
)

type IEnvService interface {
	OverwriteEnvVariables(ctx context.Context, request OverwriteEnvVariablesReq) ([]*project.EnvVariable, error)
}

type envService struct {
	envRepo        ports.IEnvRepository
	projectService services.ProjectService
}

func NewEnvService(envRepo ports.IEnvRepository, projectService services.ProjectService) IEnvService {
	return &envService{
		envRepo:        envRepo,
		projectService: projectService,
	}
}
