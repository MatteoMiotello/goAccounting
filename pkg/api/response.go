package api

type Response struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func ResFromError(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

func ResFromMessage(message string) ErrorResponse {
	return ErrorResponse{
		Error: message,
	}
}
