package projects

import (
	"context"
	"net/http"

	"pulsar/internal/core/domain/billing"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services/project"
	"time"

	"github.com/labstack/echo/v4"
)

type ProjectApiResponse struct {
	ID               string               `json:"id"`
	Name             string               `json:"name"`
	Subdomain        string               `json:"subdomain"`
	DeploymentStatus string               `json:"deploymentStatus"`
	CreatedAt        time.Time            `json:"createdAt"`
	UpdatedAt        time.Time            `json:"updatedAt"`
	PricingPlan      *billing.PricingPlan `json:"pricingPlan"`
}

func ProjectToProjectApiResponse(proj *domain.Project) ProjectApiResponse {
	return ProjectApiResponse{
		ID:               proj.ID,
		Name:             proj.Name,
		Subdomain:        proj.Subdomain,
		DeploymentStatus: string(proj.DeploymentStatus),
		CreatedAt:        proj.CreatedAt,
		UpdatedAt:        proj.UpdatedAt,
		PricingPlan:      proj.PricingPlan,
	}
}

// @Summary	Get project
// @ID			get-project
// @Accept		json
// @Produce	json
// @Success	200	{object}	ProjectApiResponse
// @Param		id	path		string	true	"project id"
// @Router		/api/projects/{id} [get]
// @Security	Bearer
// @Tags		Project
func GetProject(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		project, err := projectApi.GetProject(context.TODO(), project.GetProjectReq{ProjectId: id})

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, ProjectToProjectApiResponse(project))
	}
}
