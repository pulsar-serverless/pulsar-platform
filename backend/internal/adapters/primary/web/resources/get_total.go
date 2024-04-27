package resources

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/analytics"

	"github.com/labstack/echo/v4"
)

// @Summary		Get Total Project Resource Consumptions
// @ID			get-total-project-resource
// @Accept		json
// @Produce		json
// @Success		200		{object} any
// @Router		/api/projects/{projectId}/resources/total [get]
// @Param		projectId	path	string		true	"project id"
// @Security	Bearer
// @Tags		Resources
func GetProjectTotalUtil(resourceApi analytics.IResourceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request analytics.GetProjectResRequest

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		resources, err := resourceApi.GetTotalProjectResourceUtil(context.TODO(), request.ProjectId)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, resources)
	}
}
