package project

import (
	"context"
	domain "pulsar/internal/core/domain/project"
)

func (projectService *ProjectService) InstallProject(ctx context.Context, project *domain.Project) error {
	buildContext, err := projectService.fileRepo.CreateBuildContext(project)
	if err != nil {
		return err
	}

	containerId, err := projectService.containerService.DeployContainer(ctx, project, buildContext)
	if err != nil {
		return err
	}

	projectService.projectRepo.UpdateProject(ctx,
		project.ID,
		&domain.Project{
			ContainerId: containerId,
		},
	)

	return err
}
