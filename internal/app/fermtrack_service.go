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
	GetFermentationByID(ctx context.Context, uuid string) (*repository.Fermentation, error)
	GetFermentationAdvice(ctx context.Context, question string) (string, error)
	// CreateFermentation(f *repository.Fermentation) error
	// UpdateFermentation(f *repository.Fermentation) error
	// DeleteFermentation(uuid string) error
}

func NewFermentationService(repo repository.FermentationRepository, aiClent ai.AIClient) *FermentationService {
	return &FermentationService{repo: repo, aiClient: aiClent}
}

func (s *FermentationService) GetFermentations(ctx context.Context) ([]repository.Fermentation, error) {
	return s.repo.FindAll()
}

func (s *FermentationService) GetFermentationByID(ctx context.Context, uuid string) (*repository.Fermentation, error) {
	return s.repo.FindByID(uuid)
}

// TODO join AI answer
// TODO test
// TODO moe repo and ai into internal/app/repository and internal/app/aiclient
func (s *FermentationService) GetFermentationAdvice(ctx context.Context, question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask me, the wine wizard, anything you like.", nil
	}
	// 	time.Sleep(1 * time.Second) // simulate slow AI answer
	// 	return `My apologies, but at the moment, my thoughts seem to be a bit hazy, much like a foggy morning.

	// It's as if my mind is a glass of wine, swirling with ideas, but not quite focused enough to give you a clear answer.

	// Perhaps we could revisit this later when my mind is a bit more sober and my thoughts are clearer.

	// Enjoy this poem instead:

	// Fermenting mind swirls,
	// Grape juice turned to liquid fire,
	// Drunkard dreams of wine.
	// 	`, nil
	result, err := s.aiClient.AskQuestion(ctx, question) // TODO passin context
	if err != nil {
		return "", err
	}
	return result, nil
}

// func (s *FermentationService) CreateFermentation(f *repository.Fermentation) error {
// 	if f.Nickname == "" {
// 		return errors.New("nickname cannot be empty")
// 	}
// 	return s.repo.Create(f)
// }

// func (s *FermentationService) UpdateFermentation(f *repository.Fermentation) error {
// 	return s.repo.Update(f)
// }

// func (s *FermentationService) DeleteFermentation(uuid string) error {
// 	return s.repo.Delete(uuid)
// }
