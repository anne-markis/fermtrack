package handlers

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anne-markis/fermtrack/internal/app/mocks"
	"github.com/anne-markis/fermtrack/internal/app/repository"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetFermentationsHandler(t *testing.T) {
	mockService := new(mocks.FermentationTrackService)
	handler := NewFermentationHandler(mockService)

	uuid1 := uuid.NewString()

	mockFermentations := []repository.Fermentation{
		{UUID: uuid1, Nickname: "Test Fermentation"},
	}

	mockService.EXPECT().GetFermentations(mock.Anything).Return(mockFermentations, nil)

	req, _ := http.NewRequest("GET", "/fermentations", nil)
	rr := httptest.NewRecorder()

	handler.GetFermentations(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fermentations []repository.Fermentation
	err := json.Unmarshal(rr.Body.Bytes(), &fermentations)
	require.Nil(t, err)
	assert.Len(t, fermentations, 1)
	assert.Equal(t, uuid1, fermentations[0].UUID)

	mockService.AssertExpectations(t)
}

func TestGetFermentationHandler(t *testing.T) {
	mockService := new(mocks.FermentationTrackService)
	handler := NewFermentationHandler(mockService)

	uuid1 := "123"
	mockFermentation := &repository.Fermentation{UUID: uuid1, Nickname: "Test Fermentation"}
	mockService.EXPECT().GetFermentationByUUID(mock.Anything, uuid1).Return(mockFermentation, nil)

	req, _ := http.NewRequest("GET", "/fermentations/123", nil)
	rr := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"uuid": uuid1})

	handler.GetFermentation(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fermentation repository.Fermentation
	err := json.Unmarshal(rr.Body.Bytes(), &fermentation)
	require.Nil(t, err)
	assert.Equal(t, "Test Fermentation", fermentation.Nickname)

	mockService.AssertExpectations(t)
}
