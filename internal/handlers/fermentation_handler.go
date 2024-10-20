package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/anne-markis/fermtrack/internal/app/domain"

	"github.com/gorilla/mux"
)

type FermentationHandler struct {
	service app.FermentationTrackService
}

func NewFermentationHandler(service app.FermentationTrackService) *FermentationHandler {
	return &FermentationHandler{service}
}

// GetFermentations gets a list of all fermentations
func (h *FermentationHandler) GetFermentations(w http.ResponseWriter, r *http.Request) {
	fermentations, err := h.service.GetFermentations(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fermentations)
}

// GetFermentation gets information on a particular fermentation
func (h *FermentationHandler) GetFermentation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fermentation, err := h.service.GetFermentationByUUID(r.Context(), uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(fermentation)
}

func (h *FermentationHandler) UpdateFermentation(w http.ResponseWriter, r *http.Request) {
	var fermentation domain.Fermentation
	if err := json.NewDecoder(r.Body).Decode(&fermentation); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	if err := h.service.UpdateFermentation(r.Context(), uuid, fermentation); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type FermentationQuestion struct {
	Question string `json:"question"`
}
type FermentationAdvice struct {
	Answer string `json:"answer"`
}

// GetFermentationAdvice passes a generic question to an LLM and gets a response
func (h *FermentationHandler) GetFermentationAdvice(w http.ResponseWriter, r *http.Request) {
	var question FermentationQuestion
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.service.GetFermentationAdvice(r.Context(), question.Question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	answer := FermentationAdvice{
		Answer: result,
	}
	json.NewEncoder(w).Encode(answer)
}
