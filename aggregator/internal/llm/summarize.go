package llm

import (
	"context"
	"os"
	"sync"
	openai "github.com/sashabaranov/go-openai"
)

var (
	client *openai.Client
	clientOnce sync.Once
)

func getClient() *openai.Client {
	clientOnce.Do(func() {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			panic("OPENAI_API_KEY environment variable is not set")
		}
		client = openai.NewClient(apiKey)
	})
	return client
}

func Summarize(content string) (string, error) {
	resp, err := getClient().CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo1106,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:	"system",
					Content: "You are a helpful assistant that summarizes articles for daily audio briefings in one sentence.",
				},
				{
					Role:   "user",
					Content: "Please summarize the following article content into one sentence:\n\n" + content,
				},
			},
			Temperature: 0.7,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
