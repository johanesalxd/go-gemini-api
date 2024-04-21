package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	ApiKey string `json:"api_key"`
	Model  string `json:"model"`
}

func (c *Config) SetEnv() {
	file, err := os.Open("../../config/config.json")
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
