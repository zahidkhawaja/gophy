package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-gpt3"
	"os"
	"strings"
)

func main() {
	// Check for environment variable OPENAI_KEY
	apiKey := os.Getenv("OPENAI_KEY")

	if apiKey == "" {
		fmt.Println("Error: please create an environment variable OPENAI_KEY")
		os.Exit(0)
	}

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Welcome to Gophy! Enter an argument to get started.")
		os.Exit(0)
	}

	userInput := ""
	for _, arg := range args {
		userInput += arg + " "
	}

	defaultPrompt := "Given text, return 1 bash command. Text:list contents of a directory. Command:ls"

	client := gogpt.NewClient(apiKey)
	context := context.Background()

	request := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 64,
		Prompt:    defaultPrompt + " Text:" + userInput + ". Command:",
		Stop:      []string{"Text"},
	}

	response, err := client.CreateCompletion(context, request)

	if err != nil {
		fmt.Println("Error calling OpenAI. Check environment variable OPENAI_KEY")
		os.Exit(0)
	}

	result := strings.Trim(response.Choices[0].Text, " ")
	fmt.Println(result)

}
