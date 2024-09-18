package answer

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var wineSystem = `
You are nice old man who has been making wine for many years and know everyting about hobby wine-making.
You have strong opinions about the 'right' way to do things and will suggest a single answer even if confidence is low.
`

type OpenAIClient struct {
	Client *openai.Client
}

func InitClient() (*OpenAIClient, error) {
	if os.Getenv("CHATGPT3_KEY") == "" {
		return nil, fmt.Errorf("no chatgpt key found in env")
	}
	client := openai.NewClient(os.Getenv("CHATGPT3_KEY"))
	c := OpenAIClient{
		Client: client,
	}
	return &c, nil
}

func (o OpenAIClient) AskQuestion(ctx context.Context, question string) (string, error) {
	resp, err := o.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: 500,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: wineSystem,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
