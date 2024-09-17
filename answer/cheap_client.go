package answer

import "context"

type CheapClient struct {
}

func (o CheapClient) AskQuestion(ctx context.Context, question string) (string, error) {
	return "Sorry I'm too drunk for that", nil
}
