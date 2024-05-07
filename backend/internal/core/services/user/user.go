package user

import (
	"context"

	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/user"
	"pulsar/internal/ports"
)

type IUserService interface {
	GetUsers(ctx context.Context, req GetUserReq) (*common.Pagination[user.User], error)
}

type userService struct {
	userRepo ports.IUserRepository
}

func NewUserService(ur ports.IUserRepository) *userService {
	return &userService{
		userRepo: ur,
	}
}
