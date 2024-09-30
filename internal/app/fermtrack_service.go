//go:generate mockery --name=FermentationTrackService --dir=internal/app --output=internal/mocks --with-expecter
package app

import (
	"strings"
	"time"

	"github.com/anne-markis/fermtrack/internal/repository"
)

type FermentationService struct {
	repo repository.FermentationRepository
}

type FermentationTrackService interface {
	GetFermentations() ([]repository.Fermentation, error)
	GetFermentationByID(uuid string) (*repository.Fermentation, error)
	GetFermentationAdvice(question string) (string, error)
	// CreateFermentation(f *repository.Fermentation) error
	// UpdateFermentation(f *repository.Fermentation) error
	// DeleteFermentation(uuid string) error
}

func NewFermentationService(repo repository.FermentationRepository) *FermentationService {
	return &FermentationService{repo}
}

func (s *FermentationService) GetFermentations() ([]repository.Fermentation, error) {
	return s.repo.FindAll()
}

func (s *FermentationService) GetFermentationByID(uuid string) (*repository.Fermentation, error) {
	return s.repo.FindByID(uuid)
}

// TODO join AI answer
// TODO test
func (s *FermentationService) GetFermentationAdvice(question string) (string, error) {
	question = strings.Join(strings.Fields(question), "")
	if question == "" {
		return "Ask me, the wine wizard, anything you like.", nil
	}
	time.Sleep(1 * time.Second) // simulate slow AI answer
	return `My apologies, but at the moment, my thoughts seem to be a bit hazy, much like a foggy morning.

It's as if my mind is a glass of wine, swirling with ideas, but not quite focused enough to give you a clear answer.

Perhaps we could revisit this later when my mind is a bit more sober and my thoughts are clearer.

Enjoy this poem instead:

Fermenting mind swirls,
Grape juice turned to liquid fire,
Drunkard dreams of wine.
	`, nil
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
