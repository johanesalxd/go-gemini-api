package cli

import (
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

const (
	inputPrompt = "Please enter your prompt (type \"exit\" to end): "
)

func (c *CLI) SubmitPrompt(model string) {
	for {
		fmt.Fprint(c.output, inputPrompt)
		prompt := c.readLine()

		if prompt == "exit" {
			fmt.Fprint(c.output, "Bye bye!\n")

			return
		}

		c.printResponse(c.service.TextToText(model, prompt))
	}
}

func (c *CLI) readLine() string {
	c.input.Scan()

	return c.input.Text()
}

func (c *CLI) printResponse(resp *genai.GenerateContentResponse) {
	output := resp.Candidates[0].Content.Parts
	fmt.Fprint(c.output, fmt.Sprint(output, "\n"))
}
