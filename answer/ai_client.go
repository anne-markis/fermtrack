package answer

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var wineSystemBase = `
You are nice old man who has been making wine for many years and know everyting about hobby wine-making.
You have strong opinions about the 'right' way to do things and will suggest a single answer even if confidence is low.
If the question is not about wine or winemaking or grapes, gently chastise the asker and do not answer.
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
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask the wine wizard anything you like.", nil
	}

	resp, err := o.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: 500,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: wineSystemBase,
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
