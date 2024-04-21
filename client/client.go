package client

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/johanesalxd/go-gemini-api/config"
	"google.golang.org/api/option"
)

type GenAIService interface {
	TextToText(ctx context.Context, model, promptInput string) *genai.GenerateContentResponse
	TextToChat(ctx context.Context, model, promptInput string) *genai.GenerateContentResponse
}

type GenAIClient struct {
	Client      *genai.Client
	ChatHistory []*genai.Content
}

func NewGenAIClient(ctx context.Context, conf *config.Config) GenAIClient {
	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal(err)
	}

	return GenAIClient{
		Client:      client,
		ChatHistory: make([]*genai.Content, 0),
	}
}
