package apps

import (
	"context"
	"fmt"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/adapters/primary/web/utils"
	"pulsar/internal/core/domain/analytics"
	services "pulsar/internal/core/services/analytics"
	"pulsar/internal/core/services/billing"
	"pulsar/internal/core/services/container"
	"pulsar/internal/core/services/project"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/labstack/echo/v4"
)

// @Summary	test serverless function
// @ID			exec-app
// @Router		/ [get]
// @Tags		App
func ExecuteFunction(
	containerService container.IContainerService,
	projectService project.IProjectService,
	analyticsService services.IAnalyticsService,
	billingService billing.IBillingService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		successChan := make(chan *nat.PortBinding, 1)
		errorChan := make(chan error, 1)

		subdomain := utils.GetSubdomain(ctx.Request().Host)
		project, err := projectService.GetProjectByDomain(context.Background(), project.GetProjectReq{Subdomain: subdomain})

		if err != nil {
			resp := apierrors.FromError(err)
			return ctx.JSON(resp.Status, resp)
		}

		fmt.Println(project.PricingPlan)

		if project.PricingPlan != nil {
			if err := billingService.CheckPlanLimit(context.TODO(), project.ID, project.PlanId.String()); err != nil {
				resp := apierrors.FromError(err)
				return ctx.JSON(resp.Status, resp)
			}
		}

		startTime := time.Now()
		status := analytics.Success
		containerService.StartApp(project, successChan, errorChan)

		select {
		case info := <-successChan:
			err = echo.WrapHandler(NewProxy(info))(ctx)
		case err = <-errorChan:
			status = analytics.Error
		}

		endTime := time.Now()

		go analyticsService.PublishInvocationCreatedEvent(
			context.Background(),
			analytics.New(project.ID, startTime, endTime, status))

		resp := apierrors.FromError(err)
		return ctx.JSON(resp.Status, resp)
	}
}
