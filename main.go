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

	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text(conf.PromptInput))
	if err != nil {
		log.Fatal(err)
	}

	marshalResponse, _ := json.MarshalIndent(resp, "", "	")
	fmt.Println(string(marshalResponse))
}
