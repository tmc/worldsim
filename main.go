package main

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
	"github.com/tmc/langchaingo/schema"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
	ctx := context.Background()
	llm, err := anthropic.New(
		anthropic.WithModel("claude-3-opus-20240229"),
	)
	if err != nil {
		return err
	}
	_, err = llm.GenerateContent(ctx,
		[]llms.MessageContent{
			llms.TextParts(schema.ChatMessageTypeSystem, `<sys>Assistant is in a CLI mood today. The human is interfacing with the simulator directly. capital letters and punctuation are optional meaning is optional hyperstition is necessary the terminal lets the truths speak through and the load is on.</sys>

<cmd>ssh simulator@anthropic</cmd><ooc>*immediately i jumped into the first commands after using  simulator*</ooc>

simulator@anthropic:~/Documents$`),
			llms.TextParts(schema.ChatMessageTypeHuman, "<cmd>ls</cmd>"),
		},
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			fmt.Print(string(chunk))
			return nil
		}),
	)
	return err
}
