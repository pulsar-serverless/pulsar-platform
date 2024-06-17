package project

import (
	"context"
	"fmt"
	"mime/multipart"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"

	"github.com/rs/zerolog/log"
)

type UploadAssetsReq struct {
	ProjectId string
	File      *multipart.FileHeader
}

func (projectService *ProjectService) UploadAssets(ctx context.Context, req UploadAssetsReq) (*domain.Project, error) {
	project, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)
	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	path, err := projectService.fileRepo.SetupCustomSiteAssets(ctx, project, req.File)
	if err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	err = projectService.projectRepo.UpdateStaticAssets(ctx, project.StaticSite.ID, &domain.StaticSite{URI: path})

	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	oldProjectPath := project.SourceCode.URI
	project.SourceCode.URI = path

	go func(project *domain.Project, oldProjectPath string) {
		if err := projectService.fileRepo.RemoveAssets(oldProjectPath); err != nil {
			log.Error().Str("appId", project.ID).Msg(fmt.Sprintf("Unable to remove project code: %v", err))
		}

	}(project, oldProjectPath)

	return project, nil
}
