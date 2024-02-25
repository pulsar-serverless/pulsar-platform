package projects

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

type UploadProjectCodeInput struct {
	ProjectId string `form:"projectId"`
}

// @Summary	Upload project
// @ID			upload-project-code
// @Accept		json
// @Produce	json
// @Success	200
// @Param		projectId	path	string	true	"project id"
// @Param file formData file true "zipped code"
// @Router		/api/projects/code/{projectId} [put]
// @Security	Bearer
// @Tags		Project
func UploadProjectCode(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UploadProjectCodeInput

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		file, err := c.FormFile("file")
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		project, err := projectApi.UploadProjectCode(context.TODO(), project.UploadProjectCodeReq{ProjectId: input.ProjectId, File: file})
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, project)
	}
}
