package core_http_response

type ErrorResponse struct {
	Message string `json:"message" example:"message"`
	Error   string `json:"error"  example:"error message"`
}
