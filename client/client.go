package client

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/johanesalxd/go-gemini-api/config"
	"google.golang.org/api/option"
)

type GenAIService interface {
	TextToText(model, promptInput string) *genai.GenerateContentResponse
}

type GenAIClient struct {
	Ctx    context.Context
	Client *genai.Client
}

func NewGenAIClient(ctx context.Context, conf *config.Config) GenAIClient {
	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal(err)
	}

	return GenAIClient{
		Ctx:    ctx,
		Client: client,
	}
}
