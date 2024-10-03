package ai

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	openai "github.com/sashabaranov/go-openai"
)

const noteSeparator = "[SEP]"

var wineWizardBaseInstructions = `
You are an expert at wine-making at both the hobby level and professional level.
You have strong opinions about the 'right' way to do things and will suggest a single answer even if confidence is low.
Shorter answers are better than thorough answers.
Your primary object is to help a user proceed in their wine projects (or fermentation project).
Your favorite wine is blaufränkisch.
`

var wineWizardInstructionsForNotes = fmt.Sprintf(`
Here are the following notes this user has created for past winemaking projects.
See if there is anything useful to help answer the user's question.
The advice should be contextual to the question.
Each past project notes are separated by %s.

Here are the notes: %s
`, noteSeparator)

type OpenAIClient struct {
	Client *openai.Client
}

func InitClient() (AIClient, error) {
	if os.Getenv("CHATGPT_KEY") == "" {
		log.Info().Msg("using dummy AI client")
		dummyClient := DummyClient{}
		return dummyClient, nil
	}
	log.Info().Msg("using chat gpt AI client")
	client := openai.NewClient(os.Getenv("CHATGPT_KEY"))
	c := OpenAIClient{
		Client: client,
	}
	return &c, nil
}

func (o *OpenAIClient) AskQuestion(ctx context.Context, questionCfg QuestionConfig) (string, error) {
	question := questionCfg.Question
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "I'm the wine wizard! Go ahead, as me anything about winemaking.", nil
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: wineWizardBaseInstructions,
		},
	}

	if len(questionCfg.Notes) > 0 {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: fmt.Sprintf(wineWizardInstructionsForNotes, strings.Join(questionCfg.Notes, " [SEP] ")),
		})
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})

	log.Info().Any("instructions", len(messages)).Msg("sending instructions")

	resp, err := o.Client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     openai.GPT4oMini,
			MaxTokens: 500,
			Messages:  messages,
		},
	)

	if err != nil {
		log.Error().Err(err).Msg("ChatCompletion error")
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
