// middleware/jwt.go

package auth

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"pulsar/internal/adapters/primary/web/apierrors"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	return func(c echo.Context) error {
		errorHandler := func(c echo.Context, message string) error {
			return c.JSON(http.StatusUnauthorized,
				apierrors.ApiError{
					Status:  http.StatusUnauthorized,
					Message: message,
				})
		}

		token, err := jwtmiddleware.AuthHeaderTokenExtractor(c.Request())
		if err != nil {
			return errorHandler(c, "Invalid authorization token format")
		}

		claims, err := jwtValidator.ValidateToken(context.Background(), token)
		if err != nil {
			return errorHandler(c, "Invalid authorization token")
		}

		userId := claims.(*validator.ValidatedClaims).RegisteredClaims.Subject
		c.Set("userId", userId)

		return next(c)
	}
}
