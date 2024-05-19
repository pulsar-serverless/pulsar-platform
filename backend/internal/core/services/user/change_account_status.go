package user

import (
	"context"
	"pulsar/internal/core/services"
)

type ChangeAccountStatusReq struct {
	UserId string `param:"id"`
	Status string `form:"status"`
}

type ChangeAccountStatusRes = string

func (projectService *userService) ChangeAccountStatus(ctx context.Context, req ChangeAccountStatusReq) error {
	err := projectService.userRepo.ChangeAccountStatus(ctx, req.UserId, req.Status)

	if err != nil {
		return services.NewAppError(services.ErrInternalServer, err)
	}
	return nil
}
