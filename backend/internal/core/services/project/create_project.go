package project

import (
	"context"
	"fmt"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"

	"github.com/google/uuid"
	"github.com/rs/xid"
)

type CreateProjectReq struct {
	ProjectName string
	UserId      string
}

type GenericProjectResp struct {
	ID               string                   `json:"id"`
	Name             string                   `json:"name"`
	DeploymentStatus project.DeploymentStatus `json:"deploymentStatus"`
	CreatedAt        time.Time                `json:"createdAt"`
	UpdatedAt        time.Time                `json:"updatedAt"`
}

func GenericProjectRespFromProject(project *project.Project) *GenericProjectResp {
	return &GenericProjectResp{
		ID:               project.ID,
		Name:             project.Name,
		DeploymentStatus: project.DeploymentStatus,
		CreatedAt:        project.CreatedAt,
		UpdatedAt:        project.UpdatedAt,
	}
}

func (projectService *ProjectService) CreateProject(ctx context.Context, req CreateProjectReq) (*GenericProjectResp, error) {
	var newProject = project.Project{
		ID:     fmt.Sprintf("%s-%s", req.ProjectName, generateAppId()),
		Name:   req.ProjectName,
		UserId: req.UserId,
	}

	if err := projectService.projectRepo.CreateProject(ctx, &newProject); err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	go func(newProject *project.Project) {
		sourceDir, err := projectService.fileRepo.SetupDefaultProject(newProject)
		if err != nil {
			return
		}

		newProject.SourceCode = &project.SourceCode{URI: sourceDir, ID: uuid.New()}

		_, err = projectService.projectRepo.UpdateProject(ctx, newProject.ID, newProject)
		if err != nil {
			return
		}

		buildContext, err := projectService.fileRepo.CreateBuildContext(newProject)
		if err != nil {
			return
		}

		projectService.containerService.DeployContainer(ctx, newProject, buildContext)
	}(&newProject)

	return GenericProjectRespFromProject(&newProject), nil
}

func generateAppId() string {
	return xid.New().String()
}
