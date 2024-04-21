package cli

import (
	"bufio"
	"io"

	"github.com/johanesalxd/go-gemini-api/server"
)

type CLI struct {
	input   *bufio.Scanner
	output  io.Writer
	service server.GenAIService
}

func NewCLI(input io.Reader, output io.Writer, service server.GenAIService) *CLI {
	return &CLI{
		input:   bufio.NewScanner(input),
		output:  output,
		service: service,
	}
}
