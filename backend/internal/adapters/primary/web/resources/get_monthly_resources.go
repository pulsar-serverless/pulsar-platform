package resources

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/analytics"

	"github.com/labstack/echo/v4"
)

// @Summary	Get Monthly Project Resource Consumptions
// @ID			get-monthly-project-resource
// @Accept		json
// @Produce	json
// @Success	200		{object} any
// @Router		/api/projects/{projectId}/resources/monthly [get]
// @Param		projectId		path		string					true	"project id"
// @Param		month	query		string	false	"Month"
// @Security	Bearer
// @Tags		Resources

func GetProjectMonthlyConsumption(resourceApi analytics.IResourceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request analytics.GetProjectResRequest

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		resources, err := resourceApi.GetMonthlyProjectResourceUtil(context.TODO(), request.ProjectId, request.Month)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, resources)
	}
}
