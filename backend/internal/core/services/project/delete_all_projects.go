package project

import (
	"context"
	"fmt"
	"pulsar/internal/core/services"
)

type DeleteAllProjectsReq struct {
	UserId string `param:"id"`
}

func (projectService *ProjectService) DeleteAllProjects(ctx context.Context, req DeleteAllProjectsReq) error {
	projects, err := projectService.projectRepo.GetAllProjects(ctx, req.UserId)
	if err != nil {
		return services.NewAppError(services.ErrNotFound, err)
	}

	for _, project := range projects {
		fmt.Println(project.UserId)
		if err := projectService.DeleteProject(ctx, DeleteProjectReq{ProjectId: project.ID}); err != nil {
			return err
		}
	}

	return nil
}
