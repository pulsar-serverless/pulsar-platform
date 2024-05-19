package users

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/user"

	"github.com/labstack/echo/v4"
)

// @Summary	Get user account status
// @ID			get-users-account-status
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]any
// @Router		/api/users/status [get]
// @Security	Bearer
// @Tags		USER
func GetAccountStatus(userApi user.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		status, err := userApi.GetUserStatus(context.TODO(), c.Get("userId").(string))
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.JSON(http.StatusOK, &status)
	}
}
