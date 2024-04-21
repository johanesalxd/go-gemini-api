package main

import (
	"context"
	"os"

	"github.com/johanesalxd/go-gemini-api/cli"
	"github.com/johanesalxd/go-gemini-api/client"
	"github.com/johanesalxd/go-gemini-api/config"
)

func main() {
	conf := new(config.Config)
	conf.GetEnv()

	genAI := client.NewGenAIClient(context.Background(), conf)
	defer genAI.Client.Close()

	cli := cli.NewCLI(os.Stdin, os.Stdout, &genAI)
	cli.SubmitPrompt(conf.Model)
}
