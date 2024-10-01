//go:generate mockery --name=FermentationTrackService --dir=internal/app --output=internal/app/mocks --with-expecter
package app

import (
	"context"
	"strings"

	"github.com/anne-markis/fermtrack/internal/app/ai"
	"github.com/anne-markis/fermtrack/internal/app/repository"
)

type FermentationService struct {
	repo     repository.FermentationRepository
	aiClient ai.AIClient
}

type FermentationTrackService interface {
	GetFermentations(ctx context.Context) ([]repository.Fermentation, error)
	GetFermentationByUUID(ctx context.Context, uuid string) (*repository.Fermentation, error)
	GetFermentationAdvice(ctx context.Context, question string) (string, error)
}

func NewFermentationService(repo repository.FermentationRepository, aiClent ai.AIClient) *FermentationService {
	return &FermentationService{repo: repo, aiClient: aiClent}
}

func (s *FermentationService) GetFermentations(ctx context.Context) ([]repository.Fermentation, error) {
	return s.repo.FindAll()
}

func (s *FermentationService) GetFermentationByUUID(ctx context.Context, uuid string) (*repository.Fermentation, error) {
	return s.repo.FindByUUID(uuid)
}

func (s *FermentationService) GetFermentationAdvice(ctx context.Context, question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask me, the wine wizard, anything you like.", nil
	}

	// TODO: this gets all fermentations. Should only get undeleted fermentations from loggedin use
	// It would also make sense to only get notes relevant to the question, (filtered by type of wine?)
	pastFerms, err := s.repo.FindAll()
	if err != nil {
		return "", err
	}
	pastNotes := make([]string, len(pastFerms))
	for i, ferm := range pastFerms {
		pastNotes[i] = ferm.RecipeNotes // TODO pass along
	}

	result, err := s.aiClient.AskQuestion(ctx, ai.QuestionConfig{
		Question: question,
		Notes:    pastNotes,
	})
	if err != nil {
		return "", err
	}
	return result, nil
}
