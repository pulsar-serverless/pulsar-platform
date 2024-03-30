package project

import (
	"context"
	"pulsar/internal/core/services"
)

type RemoveAPIKeyReq struct {
	ProjectId string `param:"id"`
}

func (service *ProjectService) RemoveAPIkey(ctx context.Context, request RemoveAPIKeyReq) error {
	updates := make(map[string]interface{})
	updates["TokenIssuedAt"] = nil

	if _, err := service.projectRepo.UpdateProjectFields(ctx, request.ProjectId, updates); err != nil {
		return services.NewAppError(services.ErrInternalServer, err)
	}

	return nil
}
