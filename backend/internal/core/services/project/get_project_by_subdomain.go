package project

import (
	"context"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
)

func (projectService *ProjectService) GetProjectBySubdomain(ctx context.Context, subdomain string) (*project.Project, error) {
	project, err := projectService.projectRepo.GetProjectBySubdomain(ctx, subdomain)

	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	return project, nil
}
