package envs

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	projectServices "pulsar/internal/core/services/project"

	zeroLog "github.com/rs/zerolog/log"
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
		envServices.projectService.UpdateProject(
			ctx,
			projectServices.UpdateProjectReq{
				ProjectId:      existingProject.ID,
				UpdatedProject: &project.Project{DeploymentStatus: project.Building},
			},
		)

		err := envServices.projectService.InstallProject(context.TODO(), existingProject)
		if err != nil {
			zeroLog.Error().
				Str("AppID", existingProject.ID).
				Err(err).
				Msg("Unable to install a project.")

			envServices.projectService.UpdateProject(
				ctx,
				projectServices.UpdateProjectReq{
					ProjectId:      existingProject.ID,
					UpdatedProject: &project.Project{DeploymentStatus: project.Failed},
				},
			)
			return
		}
		envServices.projectService.UpdateProject(
			ctx,
			projectServices.UpdateProjectReq{
				ProjectId:      existingProject.ID,
				UpdatedProject: &project.Project{DeploymentStatus: project.Done},
			},
		)
	}()

	return variables, nil
}
