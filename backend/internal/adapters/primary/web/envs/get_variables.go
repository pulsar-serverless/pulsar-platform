package envs

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	_ "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services/envs"

	"github.com/labstack/echo/v4"
)

// @Summary	Get Environmental Variables
// @ID			get-env-variables
// @Accept		json
// @Produce	json
// @Success	200		{object}	[]project.EnvVariable
// @Router		/api/projects/envs/{projectId} [get]
// @Param		projectId		path		string					true	"project id"
// @Security	Bearer
// @Tags		Env
func GetEnvVariables(envApi envs.IEnvService) echo.HandlerFunc {
	return func(c echo.Context) error {
		projectId := c.Param("projectId")

		envs, err := envApi.GetEnvVariables(context.TODO(), projectId)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, envs)
	}
}
