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

func (g *GenAIClient) TextToChat(ctx context.Context, model, promptInput string) *genai.GenerateContentResponse {
	mdl := g.Client.GenerativeModel(model)
	cs := mdl.StartChat()
	cs.History = append(cs.History, g.ChatHistory...)

	resp, err := cs.SendMessage(ctx, genai.Text(promptInput))
	if err != nil {
		log.Fatal(err)
	}

	g.ChatHistory = cs.History

	return resp
}
