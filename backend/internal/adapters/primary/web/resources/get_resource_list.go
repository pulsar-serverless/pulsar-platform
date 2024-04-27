package resources

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/analytics"

	"github.com/labstack/echo/v4"
)

// @Summary		Get Project Resource Consumption
// @ID			get-project-resource
// @Accept		json
// @Produce		json
// @Success		200		{object} any
// @Router		/api/projects/{projectId}/resources [get]
// @Param		projectId	path		string	true	"project id"
// @Param		pageNumber	query		int		true	"Page number"
// @Param		pageSize	query		int		true	"Page size"
// @Param		month		query		string	false	"Month"
// @Security	Bearer
// @Tags		Resources
func GetResourceUtilList(resourceApi analytics.IResourceService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request analytics.GetProjectResRequest

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		resources, err := resourceApi.GetProjectResourceUtil(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, resources)
	}
}
