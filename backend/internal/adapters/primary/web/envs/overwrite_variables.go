package envs

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/envs"

	"github.com/labstack/echo/v4"
)

// @Summary	Create Environmental Variables
// @ID			create-env-variables
// @Accept		json
// @Produce	json
// @Success	200		{object}	project.GenericProjectResp
// @Router		/api/projects/envs/{projectId} [post]
// @Param		projectId		path		string					true	"project id"
// @Param		request	body		envs.OverwriteEnvVariablesReq	true	"Create env variables"
// @Security	Bearer
// @Tags		Env
func OverwriteEnvVariables(envApi envs.IEnvService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input envs.OverwriteEnvVariablesReq
		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		envs, err := envApi.OverwriteEnvVariables(context.TODO(), input)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, envs)
	}
}
