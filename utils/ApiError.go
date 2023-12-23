package utils

type ApiError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func NewApiError(message string, errors interface{}) *ApiError {
	return &ApiError{Status: false, Message: message, Errors: errors}
}
