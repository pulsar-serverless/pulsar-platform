package user

import (
	"context"
	"pulsar/internal/core/services"
)

func (projectService *userService) GetUserStatus(ctx context.Context, userId string) (string, error) {
	status, err := projectService.userRepo.GetAccountStatus(ctx, userId)

	if err != nil {
		return "", services.NewAppError(services.ErrInternalServer, err)
	}
	return status, nil
}
