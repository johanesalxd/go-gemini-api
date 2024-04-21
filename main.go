package main

import (
	"context"
	"os"

	"github.com/johanesalxd/go-gemini-api/cli"
	"github.com/johanesalxd/go-gemini-api/config"
	"github.com/johanesalxd/go-gemini-api/server"
)

func main() {
	conf := new(config.Config)
	conf.GetEnv()

	genAI := server.NewGenAIClient(context.Background(), conf)
	defer genAI.Client.Close()

	cli := cli.NewCLI(os.Stdin, os.Stdout, &genAI)
	cli.SubmitPrompt(conf.Model)
}
