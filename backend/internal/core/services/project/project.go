package project

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/ports"
)

type IProjectService interface {
	CreateProject(ctx context.Context, req CreateProjectReq) (*GenericProjectResp, error)
	DeleteProject(ctx context.Context, req DeleteProjectReq) error
	GetProject(ctx context.Context, req GetProjectReq) (*GenericProjectResp, error)
	GetProjects(ctx context.Context, req GetProjectsReq) (*common.Pagination[GenericProjectResp], error)
	UpdateProject(ctx context.Context, req UpdateProjectReq) (*GenericProjectResp, error)
}

type ProjectService struct {
	projectRepo ports.IProjectRepo
}

func NewProjectService(pr ports.IProjectRepo) *ProjectService {
	return &ProjectService{
		projectRepo: pr,
	}
}
