package apps

import (
	"context"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/adapters/primary/web/utils"
	"pulsar/internal/core/domain/analytics"
	services "pulsar/internal/core/services/analytics"
	"pulsar/internal/core/services/container"
	"pulsar/internal/core/services/project"
	"time"

	"github.com/labstack/echo/v4"
)

// @Summary	test serverless function
// @ID			exec-app
// @Router		/ [get]
// @Tags		App
func ExecuteFunction(
	containerService container.IContainerService,
	projectService project.IProjectService,
	analyticsService services.IAnalyticsService) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			successChan := make(chan bool, 1)
			errorChan := make(chan error, 1)

			subdomain := utils.GetSubdomain(ctx.Request().Host)
			project, err := projectService.GetProject(context.Background(), project.GetProjectReq{ProjectId: subdomain})

			if err != nil {
				resp := apierrors.FromError(err)
				return ctx.JSON(resp.Status, resp)
			}

			startTime := time.Now()
			status := analytics.Success
			containerService.StartApp(project, successChan, errorChan)

			select {
			case <-successChan:
				err = next(ctx)
			case err := <-errorChan:
				resp := apierrors.FromError(err)
				err = ctx.JSON(resp.Status, resp)
				status = analytics.Error
			}

			endTime := time.Now()

			go analyticsService.PublishInvocationCreatedEvent(
				context.Background(),
				analytics.New(project.ID, startTime, endTime, status))

			return err
		}
	}
}
