package services

import "errors"

var (
	ErrBadRequest     = errors.New("Bad request")
	ErrInternalServer = errors.New("Internal server error")
	ErrNotFound       = errors.New("Not found")
)

type AppError struct {
	serviceError  error
	externalError error
}

func (e AppError) ServiceError() error {
	return e.serviceError
}

func (e AppError) ExternalError() error {
	return e.externalError
}

func NewAppError(serviceError, externalError error) error {
	return AppError{
		serviceError:  serviceError,
		externalError: externalError,
	}
}

func (e AppError) Error() string {
	return errors.Join(e.serviceError, e.externalError).Error()
}
