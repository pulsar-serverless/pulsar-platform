package ports

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"

	"github.com/google/uuid"
)

type IProjectRepo interface {
	CreateProject(ctx context.Context, project *project.Project) error
	UpdateProject(ctx context.Context, projectId string, updatedProject *project.Project) (*project.Project, error)
	GetProject(ctx context.Context, projectId string) (*project.Project, error)
	GetProjects(ctx context.Context, pageNumber int, pageSize int, userId string) (*common.Pagination[project.Project], error)
	GetAllProjects(ctx context.Context, userId string) ([]*project.Project, error)
	DeleteProject(ctx context.Context, projectId string) error
	UpdateSourceCode(ctx context.Context, id uuid.UUID, code *project.SourceCode) error
	UpdateProjectFields(ctx context.Context, projectId string, updatedProject map[string]interface{}) (*project.Project, error)
}
