package project

import (
	"context"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"
	"time"

	"github.com/golang-jwt/jwt"
)

func generateJWTToken(secreteKey string, iat *time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["tiat"] = iat

	return token.SignedString([]byte(secreteKey))
}

type GenerateAPITokenReq struct {
	ProjectId string `param:"id"`
}

type GenerateAPITokenRes struct {
	Token string `json:"token"`
}

func (service *ProjectService) GenerateAPIToken(ctx context.Context, request GenerateAPITokenReq) (*GenerateAPITokenRes, error) {
	iat := time.Now()

	token, err := generateJWTToken(service.jwtSecreteKey, &iat)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	_, err = service.projectRepo.UpdateProject(ctx,
		request.ProjectId,
		&domain.Project{TokenIssuedAt: &iat},
	)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return &GenerateAPITokenRes{token}, nil
}
