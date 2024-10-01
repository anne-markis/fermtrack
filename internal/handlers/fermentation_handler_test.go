package handlers

import (
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anne-markis/fermtrack/internal/app/mocks"
	"github.com/anne-markis/fermtrack/internal/app/repository"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetFermentationsHandler(t *testing.T) {
	mockService := new(mocks.FermentationTrackService)
	handler := NewFermentationHandler(mockService)

	mockFermentations := []repository.Fermentation{
		{UUID: "123", Nickname: "Test Fermentation"},
	}

	mockService.On("GetFermentations").Return(mockFermentations, nil)

	req, _ := http.NewRequest("GET", "/fermentations", nil)
	rr := httptest.NewRecorder()

	handler.GetFermentations(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fermentations []repository.Fermentation
	err := json.Unmarshal(rr.Body.Bytes(), &fermentations)
	assert.Nil(t, err)
	assert.Equal(t, "123", fermentations[0].UUID)

	mockService.AssertExpectations(t)
}

func TestGetFermentationHandler(t *testing.T) {
	mockService := new(mocks.FermentationTrackService)
	handler := NewFermentationHandler(mockService)

	mockFermentation := &repository.Fermentation{UUID: "123", Nickname: "Test Fermentation"}
	mockService.On("GetFermentationByID", "123").Return(mockFermentation, nil) // TODO not this style

	req, _ := http.NewRequest("GET", "/fermentations/123", nil)
	rr := httptest.NewRecorder()

	// Add the URL variables to the request
	req = mux.SetURLVars(req, map[string]string{"uuid": "123"})

	handler.GetFermentation(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var fermentation repository.Fermentation
	err := json.Unmarshal(rr.Body.Bytes(), &fermentation)
	assert.Nil(t, err)
	assert.Equal(t, "Test Fermentation", fermentation.Nickname)

	mockService.AssertExpectations(t)
}
