package client

import (
	"log"

	"github.com/google/generative-ai-go/genai"
)

func (g *GenAIClient) TextToText(model, promptInput string) *genai.GenerateContentResponse {
	mdl := g.Client.GenerativeModel(model)

	resp, err := mdl.GenerateContent(g.Ctx, genai.Text(promptInput))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
