package project

import (
	"context"
	"pulsar/internal/core/services"
)

type GetProjectReq struct {
	ProjectId string
}

func (projectService *ProjectService) GetProject(ctx context.Context, req GetProjectReq) (*GenericProjectResp, error) {
	project, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)

	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	return GenericProjectRespFromProject(project), nil
}
