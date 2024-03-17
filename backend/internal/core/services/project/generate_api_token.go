package project

import (
	"context"
	"crypto/rand"
	b64 "encoding/base64"
	domain "pulsar/internal/core/domain/project"
	"pulsar/internal/core/services"

	"golang.org/x/crypto/bcrypt"
)

const TOKEN_LENGTH = 32

func generateKey(length int) (string, error) {
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return b64.StdEncoding.EncodeToString(key), nil
}

type GenerateAPITokenReq struct {
	ProjectId string `param:"id"`
}

type GenerateAPITokenRes struct {
	Token string `json:"token"`
}

func (service *ProjectService) GenerateAPIToken(ctx context.Context, request GenerateAPITokenReq) (*GenerateAPITokenRes, error) {
	token, err := generateKey(TOKEN_LENGTH)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), 14)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	_, err = service.projectRepo.UpdateProject(ctx,
		request.ProjectId,
		&domain.Project{
			ApiToken: string(hashedToken),
		},
	)
	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return &GenerateAPITokenRes{token}, nil
}
