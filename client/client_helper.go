package client

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
)

func (g *GenAIClient) TextToText(ctx context.Context, model, promptInput string) *genai.GenerateContentResponse {
	mdl := g.Client.GenerativeModel(model)

	resp, err := mdl.GenerateContent(ctx, genai.Text(promptInput))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
