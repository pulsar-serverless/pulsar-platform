package log

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	_ "pulsar/internal/core/domain/common"
	"pulsar/internal/core/services/log"

	"github.com/labstack/echo/v4"
)

// @Summary	Get Project log
// @ID			get-project-log
// @Accept		json
// @Produce	json
// @Success	200		{object} any
// @Router		/api/projects/logs/{projectId} [get]
// @Param		projectId		path		string					true	"project id"
// @Param		pageNumber	query		int	true	"Page number"
// @Param		pageSize	query		int	true	"Page size"
// @Security	Bearer
// @Tags		Log
func GetProjectLogs(envApi log.ILogService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request log.GetLogsRequest

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		logs, err := envApi.GetProjectLogs(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, logs)
	}
}
