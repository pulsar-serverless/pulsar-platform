package ports

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/user"
)

type IUserRepository interface {
	GetUsers(ctx context.Context, pageSize, pageNumber int, searchQuery string) (*common.Pagination[user.User], error)
	ChangeAccountStatus(ctx context.Context, userId, status string) error
	GetAccountStatus(ctx context.Context, userId string) (string, error)
}
