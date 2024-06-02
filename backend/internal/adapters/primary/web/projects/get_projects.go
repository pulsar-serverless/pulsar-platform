package projects

import (
	"context"
	"net/http"
	"net/url"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	Get projects
// @ID			get-projects
// @Accept		json
// @Produce	json
// @Success	200			{object}	project.GenericProjectResp
// @Param		pageNumber	query		int	true	"Page number"
// @Param		pageSize	query		int	true	"Page size"
// @Router		/api/projects [get]
// @Security	Bearer
// @Tags		Project
func GetProjects(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input struct {
			PageNumber int    `query:"pageNumber"`
			PageSize   int    `query:"pageSize"`
			UserId     string `query:"userId"`
		}

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		if input.UserId == "" {
			input.UserId = c.Get("userId").(string)
		} else {
			decoded, err := url.PathUnescape(input.UserId)
			if err == nil {
				input.UserId = string(decoded)
			}
		}

		request := project.GetProjectsReq{PageNumber: input.PageNumber, PageSize: input.PageSize, UserId: input.UserId}
		projects, err := projectApi.GetProjects(context.TODO(), request)

		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, projects)
	}
}
