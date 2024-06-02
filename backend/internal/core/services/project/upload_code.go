package project

import (
	"context"
	"fmt"
	"mime/multipart"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"

	"github.com/rs/zerolog/log"
)

type UploadProjectCodeReq struct {
	ProjectId string
	File      *multipart.FileHeader
}

func (projectService *ProjectService) UploadProjectCode(ctx context.Context, req UploadProjectCodeReq) (*domain.Project, error) {

	existingProject, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)
	if err != nil {
		return nil, services.NewAppError(services.ErrNotFound, err)
	}

	path, err := projectService.fileRepo.SetupCustomProjectCode(ctx, existingProject, req.File)
	if err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	err = projectService.projectRepo.UpdateSourceCode(ctx,
		existingProject.SourceCode.ID,
		&domain.SourceCode{URI: path},
	)

	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	oldProjectPath := existingProject.SourceCode.URI
	existingProject.SourceCode.URI = path

	go func(project *domain.Project, oldProjectPath string) {
		// remove old project code in the background
		if err := projectService.fileRepo.RemoveSourceCode(oldProjectPath); err != nil {
			log.Error().Str("appId", project.ID).Msg(fmt.Sprintf("Unable to remove project code: %v", err))
		}

		projectService.projectRepo.UpdateProject(ctx, project.ID, &domain.Project{DeploymentStatus: domain.Building})
		if err = projectService.InstallProject(context.TODO(), project); err != nil {
			projectService.projectRepo.UpdateProject(ctx, project.ID, &domain.Project{DeploymentStatus: domain.Failed})
			return
		}
		projectService.projectRepo.UpdateProject(ctx, project.ID, &domain.Project{DeploymentStatus: domain.Done})
	}(existingProject, oldProjectPath)

	return existingProject, nil
}
