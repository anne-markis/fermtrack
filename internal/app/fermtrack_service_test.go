package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/anne-markis/fermtrack/internal/app/mocks"
	"github.com/anne-markis/fermtrack/internal/app/repository"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetFermentations(t *testing.T) {
	t.Run("get fermentations, with error", func(t *testing.T) {
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, nil)

		mockRepo.EXPECT().FindAll().Return(nil, fmt.Errorf("fail"))

		_, err := service.GetFermentations(context.Background())

		require.NotNil(t, err)
	})

	t.Run("get fermentations, some results", func(t *testing.T) {
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, nil)

		uuid1 := uuid.NewString()

		mockFermentations := []repository.Fermentation{
			{UUID: uuid1, Nickname: uuid.NewString()},
			{UUID: uuid.NewString(), Nickname: uuid.NewString()},
		}

		mockRepo.EXPECT().FindAll().Return(mockFermentations, nil)

		fermentations, err := service.GetFermentations(context.Background())

		require.Nil(t, err)
		assert.Equal(t, 2, len(fermentations))
		assert.Equal(t, uuid1, fermentations[0].UUID)
	})
}

func TestGetFermentationByUUID(t *testing.T) {
	t.Run("no fermentation", func(t *testing.T) {
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, nil)

		uuid1 := uuid.NewString()

		mockRepo.EXPECT().FindByUUID(uuid1).Return(nil, nil)

		fermentation, err := service.GetFermentationByUUID(context.Background(), uuid1)

		require.Nil(t, err)
		assert.Nil(t, fermentation)
	})

	t.Run("found fermentation", func(t *testing.T) {
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, nil)

		uuid1 := uuid.NewString()
		nickName := uuid.NewString()

		mockFermentation := &repository.Fermentation{UUID: uuid1, Nickname: nickName}

		mockRepo.EXPECT().FindByUUID(uuid1).Return(mockFermentation, nil)

		fermentation, err := service.GetFermentationByUUID(context.Background(), uuid1)

		assert.Nil(t, err)
		assert.Equal(t, nickName, fermentation.Nickname)
	})

}

func TestGetFermentationAdvice(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		mockAi := new(mocks.AIClient)
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, mockAi)

		mockRepo.EXPECT().FindAll().Return(nil, nil)
		mockAi.EXPECT().AskQuestion(mock.Anything, mock.Anything).Return("", fmt.Errorf(""))

		_, err := service.GetFermentationAdvice(context.Background(), "who")

		require.Error(t, err)
	})

	t.Run("no error", func(t *testing.T) {
		mockAi := new(mocks.AIClient)
		mockRepo := new(mocks.FermentationRepository)
		service := NewFermentationService(mockRepo, mockAi)

		mockRepo.EXPECT().FindAll().Return(nil, nil)
		mockAi.EXPECT().AskQuestion(mock.Anything, mock.Anything).Return("advice", nil)

		advice, err := service.GetFermentationAdvice(context.Background(), "who")

		require.NoError(t, err)
		assert.Equal(t, "advice", advice)
	})
}
