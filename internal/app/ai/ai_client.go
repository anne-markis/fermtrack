//go:generate mockery --name=AIClient --dir=internal/app/ai --output=internal/app/mocks --with-expecter
package ai

import "context"

type QuestionConfig struct {
	Question string
	Notes    []string
}

type AIClient interface {
	AskQuestion(ctx context.Context, question QuestionConfig) (string, error)
}
