package answer

import (
	"context"
	"strings"
	"time"
)

type CheapClient struct {
}

func (o CheapClient) AskQuestion(ctx context.Context, question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask the wine wizard anything you like.", nil
	}
	time.Sleep(1 * time.Second) // simulate slow AI answer
	return "Sorry I'm too drunk for that", nil
}
