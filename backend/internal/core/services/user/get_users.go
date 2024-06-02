package user

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/user"
	"pulsar/internal/core/services"
)

type GetUserReq struct {
	PageNumber  int    `query:"pageNumber"`
	PageSize    int    `query:"pageSize"`
	SearchQuery string `query:"searchQuery"`
}

func (projectService *userService) GetUsers(ctx context.Context, req GetUserReq) (*common.Pagination[user.User], error) {
	data, err := projectService.userRepo.GetUsers(ctx, req.PageSize, req.PageNumber, req.SearchQuery)

	if err != nil {
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}
	return data, nil
}
