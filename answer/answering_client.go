package answer

import "context"

type AnsweringClient interface {
	AskQuestion(ctx context.Context, question string) (string, error)
}
