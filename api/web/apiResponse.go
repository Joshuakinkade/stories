package web

import (
	"github.com/joshuakinkade/web-journal/errors"
)

type APIResponse struct {
	Error struct {
		Type    string `json:"type"`
		Details any    `json:"details"`
	} `json:"error", omitempty`
	Data any `json:"data", omitempty`
}

func Success(data any) APIResponse {
	return ApiResponse{
		Data: data
	}
}

func Error(err Error) struct {
	
}
