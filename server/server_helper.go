package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RunServer() {
	router := gin.Default()
	router.POST("/generatetext", s.textToText)
	router.POST("/chat", s.textToChat)

	router.Run("localhost:8080")
}

func (s *Server) textToText(c *gin.Context) {
	var newPromptRequest promptRequest

	if err := c.Bind(&newPromptRequest); err != nil {
		log.Print(err)

		return
	}

	resp := s.service.TextToText(c.Request.Context(), newPromptRequest.Model, newPromptRequest.PromptInput)
	c.IndentedJSON(http.StatusCreated, resp.Candidates[0].Content)
}

func (s *Server) textToChat(c *gin.Context) {
	var newPromptRequest promptRequest

	if err := c.Bind(&newPromptRequest); err != nil {
		log.Print(err)

		return
	}

	resp := s.service.TextToChat(c.Request.Context(), newPromptRequest.Model, newPromptRequest.PromptInput)
	c.IndentedJSON(http.StatusCreated, resp.Candidates[0].Content)
}
