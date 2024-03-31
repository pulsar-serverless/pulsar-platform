package project

import (
	"context"
	"fmt"
	"pulsar/internal/core/services"

	"github.com/rs/zerolog/log"
)

type DeleteProjectReq struct {
	ProjectId string
}

func (projectService *ProjectService) DeleteProject(ctx context.Context, req DeleteProjectReq) error {
	existingProject, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)
	if err != nil {
		return services.NewAppError(services.ErrNotFound, err)
	}

	if err := projectService.fileRepo.RemoveSourceCode(existingProject.SourceCode.URI); err != nil {
		log.Error().Str("appId", existingProject.ID).Msg(fmt.Sprintf("Unable to remove project code: %v", err))
		return services.NewAppError(services.ErrInternalServer, err)
	}

	if err := projectService.projectRepo.DeleteProject(ctx, existingProject.ID); err != nil {
		log.Error().Str("appId", existingProject.ID).Msg(fmt.Sprintf("Unable to remove project container: %v", err))
		return services.NewAppError(services.ErrInternalServer, err)
	}

	return nil
}
