package shared

func ToDefaultResponse(code int, status bool, message string) DefaultResponse {
	return DefaultResponse{
		Code:    code,
		Status:  status,
		Message: message,
	}
}

func ToErrorResponse(data DefaultResponse, errors []string) ErrorResponse {
	return ErrorResponse{
		DefaultResponse: data,
		Error:           errors,
	}
}
