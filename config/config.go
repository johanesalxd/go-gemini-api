package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	ApiKey      string `json:"api_key"`
	Env         string `json:"env"`
	PromptInput string `json:"prompt_input"`
}

func (c *Config) GetEnv() {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, c); err != nil {
		log.Fatal(err)
	}
}
