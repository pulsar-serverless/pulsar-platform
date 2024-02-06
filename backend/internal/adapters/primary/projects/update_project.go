package projects

import (
	"context"
	"net/http"
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
// @Router		/projects/{id} [put]
func UpdateProjects(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateProjectRequest

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		request := project.UpdateProjectReq{Name: input.Name, ProjectId: input.Id}
		projects, err := projectApi.UpdateProject(context.TODO(), request)

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, projects)
	}
}
