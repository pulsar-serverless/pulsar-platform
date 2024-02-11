package projects

import (
	"context"
	"net/http"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	Get project
// @ID			get-project
// @Accept		json
// @Produce	json
// @Success	200	{object}	project.GenericProjectResp
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

		return c.JSON(http.StatusOK, project)
	}
}
