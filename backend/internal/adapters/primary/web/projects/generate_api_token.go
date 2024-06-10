package projects

import (
	"context"
	"net/http"
	_ "pulsar/docs"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/project"

	"github.com/labstack/echo/v4"
)

//	@Summary	Generate API token
//	@ID			generate-api-token
//	@Accept		json
//	@Produce	json
//	@Param		id	path		string	true	"project id"
//	@Success	200	{object}	any
//	@Router		/api/projects/{id}/api-token [put]
//	@Security	Bearer
//	@Tags		Project
func GenerateAPIToken(projectApi project.IProjectService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input project.GenerateAPITokenReq

		if err := c.Bind(&input); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		result, err := projectApi.GenerateAPIToken(context.TODO(), input)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(200, result)
	}
}
