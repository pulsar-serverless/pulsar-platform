package projects

import (
	"context"
	"fmt"
	"net/http"
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
			PageNumber int `query:"pageNumber"`
			PageSize   int `query:"pageSize"`
		}

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userId := c.Get("userId").(string)

		request := project.GetProjectsReq{PageNumber: input.PageNumber, PageSize: input.PageSize, UserId: userId}
		projects, err := projectApi.GetProjects(context.TODO(), request)

		if err != nil {
			fmt.Println(err)
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, projects)
	}
}
