package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"github.com/johanesalxd/go-gemini-api/config"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	conf := new(config.Config)
	conf.GetEnv()

	client := newClient(ctx, conf)
	defer client.Close()

	printResponse(generateContent(ctx, client, "gemini-pro", conf.PromptInput))
}

func newClient(ctx context.Context, conf *config.Config) *genai.Client {
	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func generateContent(ctx context.Context, client *genai.Client, model, promptInput string) *genai.GenerateContentResponse {
	mdl := client.GenerativeModel(model)

	resp, err := mdl.GenerateContent(ctx, genai.Text(promptInput))
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func printResponse(resp *genai.GenerateContentResponse) {
	marshalResponse, _ := json.MarshalIndent(resp, "", "	")
	fmt.Println(string(marshalResponse))
}
