package apps

import (
	"context"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/container"
	"pulsar/internal/core/services/project"

	"strings"

	"github.com/labstack/echo/v4"
)

// @Summary	test serverless function
// @ID			exec-app
// @Router		/ [get]
// @Tags		App
func ExecuteFunction(containerService container.IContainerService, projectService project.IProjectService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			successChan := make(chan bool, 1)
			errorChan := make(chan error, 1)

			subdomain := getSubdomain(ctx.Request().Host)
			project, err := projectService.GetProjectBySubdomain(context.Background(), subdomain)

			if err != nil {
				resp := apierrors.FromError(err)
				return ctx.JSON(resp.Status, resp)
			}

			containerService.StartApp(project.ContainerId, successChan, errorChan)

			select {
			case <-successChan:
				return next(ctx)
			case err := <-errorChan:
				resp := apierrors.FromError(err)
				return ctx.JSON(resp.Status, resp)
			}
		}
	}
}

func getSubdomain(hostname string) string {
	parts := strings.Split(hostname, ".")

	if len(parts) >= 2 {
		return parts[0]
	}
	return ""
}
