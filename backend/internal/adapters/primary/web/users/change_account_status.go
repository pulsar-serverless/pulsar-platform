package users

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/core/services/user"

	"github.com/labstack/echo/v4"
)

//	@Summary	Change user account status
//	@ID			change-users-account-status
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	[]any
//	@Router		/api/users/{id} [put]
//	@Param		id	path	string	true	"user id"
//	@Security	Bearer
//	@Tags		USER
func ChangeAccountStatus(userApi user.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request user.ChangeAccountStatusReq

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		err := userApi.ChangeAccountStatus(context.TODO(), request)
		if err != nil {
			errResp := apierrors.FromError(err)
			return c.JSON(errResp.Status, errResp)
		}

		return c.NoContent(http.StatusOK)
	}
}
