package answer

import (
	"context"
	"strings"
)

type CheapClient struct {
}

func (o CheapClient) AskQuestion(ctx context.Context, question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "🍷🧙 Ask the wine wizard anything you like.", nil
	}

	return "Sorry I'm too drunk for that", nil
}
