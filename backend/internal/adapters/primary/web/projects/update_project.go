package projects

import (
	"context"
	"net/http"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

type UpdateProjectRequest struct {
	Name string `body:"name"`
	Id   string `param:"id"`
}

// @Summary	Update project
// @ID			update-project
// @Accept		json
// @Produce	json
// @Success	200		{object}	project.GenericProjectResp
// @Param		request	body		UpdateProjectRequest	true	"create project DTO"
// @Param		id		path		string					true	"project id"
// @Router		/api/projects/{id} [put]
// @Security	Bearer
// @Tags		Project
func UpdateProjects(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateProjectRequest

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		request := project.UpdateProjectReq{ProjectId: input.Id, UpdatedProject: &domain.Project{Name: input.Name}}
		projects, err := projectApi.UpdateProject(context.TODO(), request)

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, projects)
	}
}
