package web

type APIResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func Success(data any) APIResponse {
	return APIResponse{
		Data: data,
	}
}

func Error(err error) APIResponse {
	return APIResponse{
		Error: err.Error(),
	}
}
