package users

import (
	"context"
	"net/http"
	"pulsar/internal/core/services/user"

	"github.com/labstack/echo/v4"
)

// @Summary	Get users
// @ID			get-users
// @Accept		json
// @Produce	json
// @Success	200	{object}	[]any
// @Router		/api/users [get]
// @Param		pageNumber	query		int	true	"Page number"
// @Param		pageSize	query		int	true	"Page size"
// @Security	Bearer
// @Tags		USER
func GetUsers(userApi user.IUserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		var request user.GetUserReq

		if err := c.Bind(&request); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		projects, err := userApi.GetUsers(context.TODO(), request)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, projects)
	}
}
