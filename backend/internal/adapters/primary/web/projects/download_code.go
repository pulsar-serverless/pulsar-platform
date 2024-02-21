package projects

import (
	"context"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	Download project
// @ID			download-project
// @Accept		json
// @Produce	json
// @Success	200
// @Param		projectId	path	string	true	"project id"
// @Router		/api/projects/code/{projectId} [get]
// @Security	Bearer
// @Tags		Project
func DownloadSourceCode(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectId := c.Param("projectId")

		path, err := projectApi.DownloadProjectCode(context.TODO(), project.GetProjectReq{ProjectId: projectId})
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}
		return c.File(path)
	}
}
