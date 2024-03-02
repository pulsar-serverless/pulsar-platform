package envs

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	projectServices "pulsar/internal/core/services/project"
)

type EnvVariables struct {
	Key   string
	Value string
}

type OverwriteEnvVariablesReq struct {
	Variables []EnvVariables `form:"variables"`
	ProjectID string         `param:"projectId" json:"-"`
}

func (envServices *envService) OverwriteEnvVariables(ctx context.Context, request OverwriteEnvVariablesReq) ([]*project.EnvVariable, error) {
	input := projectServices.GetProjectReq{ProjectId: request.ProjectID}
	existingProject, err := envServices.projectService.GetProject(ctx, input)
	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	variables := make([]*project.EnvVariable, len(request.Variables))

	for index, value := range request.Variables {
		variables[index] = &project.EnvVariable{
			ProjectID: request.ProjectID,
			Key:       value.Key,
			Value:     value.Value,
		}
	}

	envs, err := envServices.envRepo.OverwriteEnvVariables(ctx, request.ProjectID, variables)
	if err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	existingProject.EnvVariables = envs

	go func() {
		envServices.projectService.InstallProject(context.TODO(), existingProject)
	}()

	return variables, nil
}
