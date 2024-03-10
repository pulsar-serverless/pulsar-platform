package log

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/log"

	"github.com/labstack/echo/v4"
)

// @Summary	Delete Project logs
// @ID			delete-project-log
// @Accept		json
// @Produce	json
// @Success	200		{object} any
// @Router		/api/projects/logs/{projectId} [delete]
// @Param		projectId		path		string					true	"project id"
// @Security	Bearer
// @Tags		Log
func DeleteProjectLogs(logApi log.ILogService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request log.DeleteLogsRequest

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err := logApi.DeleteProjectLogs(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.NoContent(http.StatusOK)
	}
}
