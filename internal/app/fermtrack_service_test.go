package app

import (
	"context"
	"testing"

	"github.com/anne-markis/fermtrack/internal/app/mocks"
	"github.com/anne-markis/fermtrack/internal/app/repository"

	"github.com/stretchr/testify/assert"
)

func TestGetFermentations(t *testing.T) {
	mockAi := new(mocks.AIClient)
	mockRepo := new(mocks.FermentationRepository)
	service := NewFermentationService(mockRepo, mockAi)

	mockFermentations := []repository.Fermentation{
		{UUID: "123", Nickname: "Fermentation 1"},
		{UUID: "456", Nickname: "Fermentation 2"},
	}

	// Define expectations
	mockRepo.On("FindAll").Return(mockFermentations, nil)

	fermentations, err := service.GetFermentations(context.Background())

	// Assertions
	assert.Nil(t, err)
	assert.Equal(t, 2, len(fermentations))
	assert.Equal(t, "123", fermentations[0].UUID)

	// Ensure that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestGetFermentationByUUID(t *testing.T) {
	mockAi := new(mocks.AIClient)
	mockRepo := new(mocks.FermentationRepository)
	service := NewFermentationService(mockRepo, mockAi)

	mockFermentation := &repository.Fermentation{UUID: "123", Nickname: "Fermentation 1"}

	mockRepo.On("FindByID", "123").Return(mockFermentation, nil)

	fermentation, err := service.GetFermentationByUUID(context.Background(), "123")

	assert.Nil(t, err)
	assert.Equal(t, "Fermentation 1", fermentation.Nickname)
	mockRepo.AssertExpectations(t)
}

// func TestCreateFermentation(t *testing.T) {
// 	mockRepo := new(mocks.FermentationRepository)
// 	service := NewFermentationService(mockRepo)

// 	mockFermentation := &repository.Fermentation{UUID: "123", Nickname: "New Fermentation"}

// 	// Successful create
// 	mockRepo.On("Create", mockFermentation).Return(nil)

// 	err := service.CreateFermentation(mockFermentation)
// 	assert.Nil(t, err)

// 	mockRepo.AssertExpectations(t)

// 	// Invalid input (empty nickname)
// 	err = service.CreateFermentation(&repository.Fermentation{UUID: "123"})
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "nickname cannot be empty", err.Error())
// }
