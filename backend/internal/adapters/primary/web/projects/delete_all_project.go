package projects

import (
	"context"
	"net/http"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete users projects
// @ID			delete-users-projects
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]any
// @Router		/api/users/{id}/projects [delete]
// @Security	Bearer
// @Tags		USER
func DeleteAllProjects(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request project.DeleteAllProjectsReq

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err := projectApi.DeleteAllProjects(context.TODO(), request)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
