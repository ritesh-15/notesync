package utils

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status bool, message string, data interface{}) *ApiResponse {
	return &ApiResponse{Status: status, Message: message, Data: data}
}
