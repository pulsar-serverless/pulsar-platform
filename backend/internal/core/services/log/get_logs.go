package log

import (
	"context"
	"pulsar/internal/core/domain/common"
	domain "pulsar/internal/core/domain/log"
	"pulsar/internal/core/services"

	zeroLog "github.com/rs/zerolog/log"
)

type GetLogsRequest struct {
	PageNumber  int      `query:"pageNumber"`
	PageSize    int      `query:"pageSize"`
	SearchQuery string   `query:"searchQuery"`
	LogTypes    []string `query:"logTypes[]"`
	ProjectId   string   `param:"projectId"`
}

func (service *logService) GetProjectLogs(ctx context.Context, request GetLogsRequest) (*common.Pagination[domain.AppLog], error) {
	logs, err := service.logRepo.GetProjectLogs(ctx, request.ProjectId, request.LogTypes, request.SearchQuery, request.PageNumber, request.PageSize)

	if err != nil {
		zeroLog.Error().Err(err).Msg("Unable to fetch project logs")
		return nil, services.NewAppError(services.ErrInternalServer, err)
	}

	return logs, nil
}
