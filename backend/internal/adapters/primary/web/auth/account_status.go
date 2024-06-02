package auth

import (
	"context"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	domain "pulsar/internal/core/domain/user"
	"pulsar/internal/core/services/user"

	"github.com/labstack/echo/v4"
)

func AuthorizeStatus(userApi user.IUserService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().Method != "GET" {
				accountStatus, err := userApi.GetUserStatus(context.TODO(), ctx.Get("userId").(string))
				if err != nil {
					errResp := apierrors.FromError(err)
					return ctx.JSON(errResp.Status, errResp)
				}

				if accountStatus == domain.Suspended {
					return ctx.JSON(http.StatusUnauthorized, apierrors.New(http.StatusUnauthorized, "Unauthorized access: Account suspended"))
				}

			}
			return next(ctx)
		}
	}
}
