package app

import (
	"errors"

	"github.com/anne-markis/fermtrack/internal/domain"
)

// FermentationService handles the business logic for fermentations
type FermentationService struct {
	repo domain.FermentationRepository
}

// NewFermentationService creates a new service with the given repository
func NewFermentationService(repo domain.FermentationRepository) *FermentationService {
	return &FermentationService{repo}
}

func (s *FermentationService) GetFermentations() ([]domain.Fermentation, error) {
	return s.repo.FindAll()
}

func (s *FermentationService) GetFermentationByID(uuid string) (*domain.Fermentation, error) {
	return s.repo.FindByID(uuid)
}

func (s *FermentationService) CreateFermentation(f *domain.Fermentation) error {
	if f.Nickname == "" {
		return errors.New("nickname cannot be empty")
	}
	return s.repo.Create(f)
}

func (s *FermentationService) UpdateFermentation(f *domain.Fermentation) error {
	return s.repo.Update(f)
}

func (s *FermentationService) DeleteFermentation(uuid string) error {
	return s.repo.Delete(uuid)
}
