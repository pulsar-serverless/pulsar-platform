package project

import (
	"context"
	"fmt"
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

	fmt.Print(project.SourceCode.URI)
	zippedCode, err := projectService.fileRepo.ZipSourceCode(project.SourceCode.URI)
	fmt.Println(zippedCode.Name())

	if err != nil {
		return "", services.NewAppError(services.ErrInternalServer, err)
	}

	return zippedCode.Name(), nil
}
