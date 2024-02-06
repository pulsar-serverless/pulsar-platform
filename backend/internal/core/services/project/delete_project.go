package project

import (
	"context"
	"pulsar/internal/core/services"
)

type DeleteProjectReq struct {
	ProjectId string
}

func (projectService *ProjectService) DeleteProject(ctx context.Context, req DeleteProjectReq) error {
	err := projectService.projectRepo.DeleteProject(ctx, req.ProjectId)
	return services.NewAppError(services.ErrInternalServer, err)
}
