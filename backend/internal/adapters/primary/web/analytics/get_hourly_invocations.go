package analytics

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/analytics"

	"github.com/labstack/echo/v4"
)

//	@Summary	Get Hourly Project Invocations
//	@ID			get-project-hourly-invocations
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	any
//	@Router		/api/projects/{projectId}/analytics/hourly [get]
//	@Param		projectId	path	string	true	"project id"
//	@Param		status		query	string	false	"Invocation Status"
//	@Security	Bearer
//	@Tags		Analytics
func GetProjectHourlyInvocations(analyticsApi analytics.IAnalyticsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request analytics.GetInvocations

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		logs, err := analyticsApi.GetHourlyInvocations(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, logs)
	}
}
