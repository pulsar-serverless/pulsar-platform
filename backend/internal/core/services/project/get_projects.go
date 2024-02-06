package project

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/services"
)

type GetProjectsReq struct {
	PageNumber, PageSize int
}

func (projectService *ProjectService) GetProjects(ctx context.Context, req GetProjectsReq) (*common.Pagination[GenericProjectResp], error) {
	projects, err := projectService.projectRepo.GetProjects(ctx, req.PageNumber, req.PageSize)

	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	response := &common.Pagination[GenericProjectResp]{
		PageSize:   projects.PageSize,
		PageNumber: projects.PageNumber,
		TotalPages: projects.TotalPages,
	}

	response.Rows = make([]*GenericProjectResp, len(projects.Rows))
	for i, item := range projects.Rows {
		response.Rows[i] = GenericProjectRespFromProject(item)
	}

	return response, nil
}
