package log

import (
	"context"
	"pulsar/internal/core/services"

	zeroLog "github.com/rs/zerolog/log"
)

type DeleteLogsRequest struct {
	ProjectId string `param:"projectId"`
}

func (service *logService) DeleteProjectLogs(ctx context.Context, request DeleteLogsRequest) error {
	err := service.logRepo.DeleteProjectLogs(ctx, request.ProjectId)

	if err != nil {
		zeroLog.Error().Err(err).Msg("Unable to fetch project logs")
		return services.NewAppError(services.ErrInternalServer, err)
	}

	return nil
}
