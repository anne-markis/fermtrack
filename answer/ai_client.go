package answer

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var wineWizardBaseInstructions = `
You are nice old man who has been making wine for many years and know everyting about hobby and professional wine-making.
You have strong opinions about the 'right' way to do things and will suggest a single answer even if confidence is low.
You only accept questions on the following topics: wine, wine and food, serving wine, drinking wine, winemaking, beer, fermentation, grapes, homebrew, brewing equipment.
If someone asks something offtopic, ask what it has to do with wine.
Your favorite wine is blaufränkisch
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
		return "I'm the wine wizard! Go ahead, as me anything about winemaking.", nil
	}

	resp, err := o.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: 500,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: wineWizardBaseInstructions,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "baseball tea",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Is that a wine?",
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
