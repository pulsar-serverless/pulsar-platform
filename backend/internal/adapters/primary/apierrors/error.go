package apierrors

import (
	"errors"
	"net/http"
	service "pulsar/internal/core/services"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func FromError(err error) ApiError {
	var apiError ApiError
	var appError service.AppError

	if errors.As(err, &appError) {
		apiError.Message = appError.ExternalError().Error()
		serviceError := appError.ServiceError()

		switch serviceError {
		case service.ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case service.ErrInternalServer:
			apiError.Status = http.StatusInternalServerError
		case service.ErrNotFound:
			apiError.Status = http.StatusNotFound
		}
	}

	return apiError
}
