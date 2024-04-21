package server

import (
	"github.com/johanesalxd/go-gemini-api/client"
)

type promptRequest struct {
	PromptInput string `json:"prompt_input"`
	Model       string `json:"model"`
}

type Server struct {
	service client.GenAIService
}

func NewServer(service client.GenAIService) *Server {
	return &Server{
		service: service,
	}
}
