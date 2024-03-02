package envs

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

func (envServices *envService) GetEnvVariables(ctx context.Context, request string) ([]*project.EnvVariable, error) {
	vars, err := envServices.envRepo.GetEnvVariables(ctx, request)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return vars, nil
}
