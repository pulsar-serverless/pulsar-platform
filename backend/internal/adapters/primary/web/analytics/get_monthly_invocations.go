package analytics

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/analytics"

	"github.com/labstack/echo/v4"
)

// @Summary	Get Monthly Project Invocations
// @ID			get-project-monthly-invocations
// @Accept		json
// @Produce	json
// @Success	200		{object} any
// @Router		/api/projects/{projectId}/analytics/monthly [get]
// @Param		projectId		path		string					true	"project id"
// @Param		status	query		string	false	"Invocation Status"
// @Security	Bearer
// @Tags		Analytics
func GetProjectMonthlyInvocations(analyticsApi analytics.IAnalyticsService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request analytics.GetInvocations

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		logs, err := analyticsApi.GetMonthlyInvocations(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, logs)
	}
}
