package user

import (
	"context"

	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/user"
	"pulsar/internal/ports"
)

type IUserService interface {
	GetUsers(ctx context.Context, req GetUserReq) (*common.Pagination[user.User], error)
	ChangeAccountStatus(ctx context.Context, req ChangeAccountStatusReq) error
	GetUserStatus(ctx context.Context, userId string) (string, error)
}

type userService struct {
	userRepo ports.IUserRepository
}

func NewUserService(ur ports.IUserRepository) *userService {
	return &userService{
		userRepo: ur,
	}
}
