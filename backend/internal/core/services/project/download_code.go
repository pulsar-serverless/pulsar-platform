package project

import (
	"context"
	"pulsar/internal/core/services"
)

type DownloadProjectCodeReq struct {
	ProjectId string
}

func (projectService *ProjectService) DownloadProjectCode(ctx context.Context, req GetProjectReq) (string, error) {
	project, err := projectService.projectRepo.GetProject(ctx, req.ProjectId)

	if err != nil {
		return "", services.NewAppError(services.ErrNotFound, err)
	}

	zippedCode, err := projectService.fileRepo.ZipSourceCode(project.SourceCode.URI)
	if err != nil {
		return "", services.NewAppError(services.ErrInternalServer, err)
	}

	return zippedCode.Name(), nil
}
