package utils

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(message string, data interface{}) *ApiResponse {
	return &ApiResponse{Status: true, Message: message, Data: data}
}
