package projects

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete project
// @ID			delete-project
// @Accept		json
// @Produce	json
// @Success	200
// @Param		id	path	string	true	"project id"
// @Router		/api/projects/{id} [delete]
// @Security	Bearer
// @Tags		Project
func DeleteProject(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := projectApi.DeleteProject(context.TODO(), project.DeleteProjectReq{ProjectId: id})

		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.NoContent(http.StatusOK)
	}
}
