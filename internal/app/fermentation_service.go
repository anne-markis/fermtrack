//go:generate mockery --name=FermentationTrackService --dir=internal/app --output=internal/mocks --with-expecter
package app

import (
	"errors"

	"github.com/anne-markis/fermtrack/internal/repository"
)

type FermentationService struct {
	repo repository.FermentationRepository
}

type FermentationTrackService interface {
	GetFermentations() ([]repository.Fermentation, error)
	GetFermentationByID(uuid string) (*repository.Fermentation, error)
	CreateFermentation(f *repository.Fermentation) error
	UpdateFermentation(f *repository.Fermentation) error
	DeleteFermentation(uuid string) error
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

func (s *FermentationService) CreateFermentation(f *repository.Fermentation) error {
	if f.Nickname == "" {
		return errors.New("nickname cannot be empty")
	}
	return s.repo.Create(f)
}

func (s *FermentationService) UpdateFermentation(f *repository.Fermentation) error {
	return s.repo.Update(f)
}

func (s *FermentationService) DeleteFermentation(uuid string) error {
	return s.repo.Delete(uuid)
}
