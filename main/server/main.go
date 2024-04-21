package main

import (
	"context"

	"github.com/johanesalxd/go-gemini-api/client"
	"github.com/johanesalxd/go-gemini-api/config"
	"github.com/johanesalxd/go-gemini-api/server"
)

func main() {
	conf := new(config.Config)
	conf.GetEnv()

	genAI := client.NewGenAIClient(context.Background(), conf)
	defer genAI.Client.Close()

	server := server.NewServer(&genAI)
	server.RunServer()
}
