package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/services"

	zeroLog "github.com/rs/zerolog/log"
)

type GetInvocations struct {
	ProjectId string                  `param:"projectId"`
	Status    domain.InvocationStatus `query:"status"`
}

func (service *analyticsService) GetHourlyInvocations(ctx context.Context, request GetInvocations) ([]*domain.InvocationCount, error) {
	data, err := service.invocationRepo.GetInvocationsOfLast24Hours(ctx, request.ProjectId, request.Status)

	if err != nil {
		zeroLog.Error().Err(err).Msg("")
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}
	return data, nil
}
