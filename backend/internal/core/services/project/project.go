package project

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services/container"
	"pulsar/internal/ports"
)

type IProjectService interface {
	CreateProject(ctx context.Context, req CreateProjectReq) (*GenericProjectResp, error)
	DeleteProject(ctx context.Context, req DeleteProjectReq) error
	GetProject(ctx context.Context, req GetProjectReq) (*project.Project, error)
	GetProjects(ctx context.Context, req GetProjectsReq) (*common.Pagination[GenericProjectResp], error)
	UpdateProject(ctx context.Context, req UpdateProjectReq) (*GenericProjectResp, error)
	UploadProjectCode(ctx context.Context, req UploadProjectCodeReq) (*project.Project, error)
	DownloadProjectCode(ctx context.Context, req GetProjectReq) (string, error)
	InstallProject(ctx context.Context, project *project.Project) error
	GenerateAPIToken(ctx context.Context, request GenerateAPITokenReq) (*GenerateAPITokenRes, error)
}

type ProjectService struct {
	projectRepo      ports.IProjectRepo
	containerService container.IContainerService
	fileRepo         ports.IFileRepository
	jwtSecreteKey    string
}

func NewProjectService(pr ports.IProjectRepo, containerService container.IContainerService, fileRepo ports.IFileRepository, jwtSecreteKey string) *ProjectService {
	return &ProjectService{
		projectRepo:      pr,
		containerService: containerService,
		fileRepo:         fileRepo,
		jwtSecreteKey:    jwtSecreteKey,
	}
}
