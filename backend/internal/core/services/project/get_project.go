package project

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

type GetProjectReq struct {
	ProjectId string
}

func (projectService *ProjectService) GetProject(ctx context.Context, req GetProjectReq) (*project.Project, error) {
	project, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)

	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	return project, nil
}
