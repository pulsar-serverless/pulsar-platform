package apps

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/adapters/primary/web/utils"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

// @Summary	test static sites
// @ID			serve-site
// @Router		/static [get]
// @Tags		Site

func ServeSite(projectService project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		subdomain := utils.GetSubdomain(c.Request().Host)
		project, err := projectService.GetProjectByDomain(context.Background(), project.GetProjectReq{Subdomain: subdomain})

		if err != nil {
			resp := apierrors.FromError(err)
			return c.JSON(resp.Status, resp)
		}

		if project.StaticSite != nil {
			return c.File(project.StaticSite.URI)
		}

		return c.NoContent(http.StatusBadRequest)
	}
}
