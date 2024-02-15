package projects

import (
	"context"
	"net/http"
	_ "pulsar/docs"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

type CreateProjectRequest struct {
	Name string `form:"name"`
}

// @Summary	Create project
// @ID			create-project
// @Accept		json
// @Produce	json
// @Param		request	body		CreateProjectRequest	true	"create project DTO"
// @Success	200		{object}	project.GenericProjectResp
// @Router		/api/projects [post]
// @Security	Bearer
// @Tags		Project
func CreateProject(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreateProjectRequest

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		result, err := projectApi.CreateProject(
			context.TODO(),
			project.CreateProjectReq{ProjectName: input.Name})

		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(200, result)
	}
}
