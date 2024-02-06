package project

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

type UpdateProjectReq struct {
	ProjectId string
	Name      string
}

func (projectService *ProjectService) UpdateProject(ctx context.Context, req UpdateProjectReq) (*GenericProjectResp, error) {
	updatedProject := project.Project{Name: req.Name}

	project, err := projectService.projectRepo.UpdateProject(ctx, req.ProjectId, &updatedProject)
	if err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	return GenericProjectRespFromProject(project), nil
}
