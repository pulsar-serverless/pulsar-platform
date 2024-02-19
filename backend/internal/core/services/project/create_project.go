package project

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"

	"github.com/google/uuid"
)

type CreateProjectReq struct {
	ProjectName string
}

type GenericProjectResp struct {
	ID               uuid.UUID                `json:"id"`
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
		ID:        uuid.New(),
		Name:      req.ProjectName,
		Subdomain: fmt.Sprintf("%s-%s", req.ProjectName, generateAppId()),
	}

	if err := projectService.projectRepo.CreateProject(ctx, &newProject); err != nil {
		return nil, services.NewAppError(services.ErrBadRequest, err)
	}

	go func(project *project.Project) {
		projectService.containerService.DeployContainerWithStarterCode(ctx, project)
	}(&newProject)

	return GenericProjectRespFromProject(&newProject), nil
}

func generateAppId() string {
	// TODO: change this
	bytes := make([]byte, 12)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)[:6]
}
