package folcon

import (
	res "github.com/msdrepo/framework/responses"
)

// New Success creates a pre-filled SuccessResponse
func SuccessResponse(statusCode int, message string, data interface{}, token interface{}) res.SuccessResponse {
	return res.SuccessResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Meta:       nil,
		Token:      token,
	}
}

// New Eroor creates a pre-filled SuccessResponse
func ErrorResponse(statusCode int, message string, data interface{}) res.ErrorResponse {
	return res.ErrorResponse{
		Status:     "error",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

// SuccessWithPagination creates a success response including pagination meta
func SuccessWithPagination(statusCode int, message string, data interface{}, page, limit int, total int64) res.SuccessResponse {
	// Total pages
	totalPages := int(total / int64(limit))
	if total%int64(limit) != 0 {
		totalPages++
	}

	return res.SuccessResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Meta: &res.PaginationMeta{
			CurrentPage: page,
			PerPage:     limit,
			TotalPages:  totalPages,
			TotalItems:  total,
		},
		Token: nil,
	}
}
