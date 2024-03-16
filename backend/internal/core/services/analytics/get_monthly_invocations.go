package analytics

import (
	"context"
	domain "pulsar/internal/core/domain/analytics"
	"pulsar/internal/core/services"

	zeroLog "github.com/rs/zerolog/log"
)

func (service *analyticsService) GetMonthlyInvocations(ctx context.Context, request GetInvocations) ([]*domain.InvocationCount, error) {
	data, err := service.invocationRepo.GetInvocationsOfLast30Days(ctx, request.ProjectId, request.Status)

	if err != nil {
		zeroLog.Error().Err(err).Msg("")
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}
	return data, nil
}
