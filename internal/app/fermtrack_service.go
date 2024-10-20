//go:generate mockery --name=FermentationTrackService --dir=internal/app --output=internal/app/mocks --with-expecter
package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/anne-markis/fermtrack/internal/app/ai"
	repository "github.com/anne-markis/fermtrack/internal/app/domain"
)

type FermentationService struct {
	repo     repository.FermentationRepository
	aiClient ai.AIClient
}

type FermentationTrackService interface {
	GetFermentations(ctx context.Context) ([]repository.Fermentation, error)
	GetFermentationByUUID(ctx context.Context, uuid string) (*repository.Fermentation, error)
	GetFermentationAdvice(ctx context.Context, question string) (string, error)
	UpdateFermentation(ctx context.Context, uuid string, ferm repository.Fermentation) error
}

func NewFermentationService(repo repository.FermentationRepository, aiClent ai.AIClient) *FermentationService {
	return &FermentationService{repo: repo, aiClient: aiClent}
}

// GetFermentations gets all fermentations
func (s *FermentationService) GetFermentations(ctx context.Context) ([]repository.Fermentation, error) {
	return s.repo.FindAll()
}

// GetFermentationByUUID gets a single fermentation by uuid
func (s *FermentationService) GetFermentationByUUID(ctx context.Context, uuid string) (*repository.Fermentation, error) {
	return s.repo.FindByUUID(uuid)
}

// GetFermentationAdvice forwards a generic question to an LLM
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
		pastNotes[i] = ferm.RecipeNotes
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

func (s *FermentationService) UpdateFermentation(ctx context.Context, uuid string, newFerm repository.Fermentation) error {
	existingFerm, err := s.repo.FindByUUID(uuid)
	if err != nil {
		return err
	}
	if existingFerm == nil || existingFerm.IsZero() {
		return fmt.Errorf("fermentation does not exist; cannot update")
	}

	if !newFerm.BottledAt.IsZero() {
		existingFerm.BottledAt = newFerm.BottledAt
	}
	if newFerm.Nickname != "" {
		existingFerm.Nickname = newFerm.Nickname
	}

	if newFerm.TastingNotes != nil {
		existingFerm.TastingNotes = newFerm.TastingNotes
	}
	return s.repo.Update(existingFerm)
}
