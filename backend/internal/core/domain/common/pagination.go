package common

type Pagination[T any] struct {
	PageSize   int   `json:"limit"`
	PageNumber int   `json:"page"`
	TotalPages int64 `json:"totalPages"`
	Rows       []*T  `json:"rows"`
}
