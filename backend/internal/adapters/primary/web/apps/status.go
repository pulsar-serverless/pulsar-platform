package apps

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/container"

	"github.com/labstack/echo/v4"
)

// @Summary	 update serverless app status
// @ID			app-status
// @Accept		json
// @Produce	json
// @Router		/app/status [post]
// @Param        subdomain   query      string  true  "App subdomain"
// @Security	Bearer
// @Tags		App
func Status(containerService container.IContainerService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		subdomain := ctx.QueryParam("subdomain")

		err := containerService.ChangeAppStatus(context.TODO(), subdomain)
		if err != nil {
			return ctx.JSON(500, apierrors.FromError(err))
		}

		return ctx.NoContent(http.StatusOK)
	}
}
