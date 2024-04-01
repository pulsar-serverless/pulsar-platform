package project

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

type UpdateProjectReq struct {
	ProjectId      string
	UpdatedProject *project.Project
}

func (projectService *ProjectService) UpdateProject(ctx context.Context, req UpdateProjectReq) (*GenericProjectResp, error) {
	project, err := projectService.projectRepo.UpdateProject(ctx, req.ProjectId, req.UpdatedProject)
	if err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	return GenericProjectRespFromProject(project), nil
}
