package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Config struct {
	ApiKey string `json:"api_key"`
	Env    string `json:"env"`
}

func main() {
	ctx := context.Background()
	conf := new(Config)
	conf.getEnv(conf)

	client, err := genai.NewClient(ctx, option.WithAPIKey(conf.ApiKey))
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
		log.Fatal(err)
	}

	marshalResponse, _ := json.Marshal(resp)
	fmt.Println(string(marshalResponse))
}

func (c *Config) getEnv(conf *Config) {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		log.Fatal(err)
	}
}
