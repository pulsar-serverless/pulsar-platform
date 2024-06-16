package auth

import (
	"context"
	"fmt"
	"net/http"
	"pulsar/internal/adapters/primary/web/apierrors"
	"pulsar/internal/adapters/primary/web/utils"
	"pulsar/internal/core/services/project"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CustomAppClaims struct {
	Iat *time.Time `json:"iat"`
	jwt.Claims
}

func IsAuthorized(projectService project.IProjectService, JWTSecreteKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			subdomain := utils.GetSubdomain(c.Request().Host)

			project, err := projectService.GetProjectByDomain(context.Background(), project.GetProjectReq{Subdomain: subdomain})
			if err != nil {
				resp := apierrors.FromError(err)
				return c.JSON(resp.Status, resp)
			}

			if project.TokenIssuedAt == nil {
				return next(c)
			}

			stringToken := c.Request().Header.Get("X-API-KEY")
			fmt.Printf("%v", stringToken)
			token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
				return []byte(JWTSecreteKey), nil
			})

			if err != nil {
				fmt.Printf("%v", err)
				return c.NoContent(http.StatusUnauthorized)
			}

			if token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)

				if ok {
					iatString := claims["tiat"].(string)
					if iat, err := time.Parse(time.RFC3339Nano, iatString); err == nil {
						if project.TokenIssuedAt.Equal(iat) {
							return next(c)
						}
					}
				}

			}

			return c.NoContent(http.StatusUnauthorized)
		}
	}
}
