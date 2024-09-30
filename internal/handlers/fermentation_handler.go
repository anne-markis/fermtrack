package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/anne-markis/fermtrack/internal/app"

	"github.com/gorilla/mux"
)

type FermentationHandler struct {
	service app.FermentationTrackService
}

func NewFermentationHandler(service app.FermentationTrackService) *FermentationHandler {
	return &FermentationHandler{service}
}

func (h *FermentationHandler) GetFermentations(w http.ResponseWriter, r *http.Request) {
	fermentations, err := h.service.GetFermentations()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fermentations)
}

func (h *FermentationHandler) GetFermentation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	fermentation, err := h.service.GetFermentationByID(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(fermentation)
}

type FermentationQuestion struct {
	Question string `json:"question"`
}
type FermentationAdvice struct {
	Answer string `json:"answer"`
}

func (h *FermentationHandler) GetFermentationAdvice(w http.ResponseWriter, r *http.Request) {
	var question FermentationQuestion
	if err := json.NewDecoder(r.Body).Decode(&question); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := h.service.GetFermentationAdvice(question.Question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	answer := FermentationAdvice{
		Answer: result,
	}
	json.NewEncoder(w).Encode(answer)
}

// func (h *FermentationHandler) CreateFermentation(w http.ResponseWriter, r *http.Request) {
// 	var fermentation repository.Fermentation
// 	if err := json.NewDecoder(r.Body).Decode(&fermentation); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if err := h.service.CreateFermentation(&fermentation); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// func (h *FermentationHandler) UpdateFermentation(w http.ResponseWriter, r *http.Request) {
// 	var fermentation repository.Fermentation
// 	if err := json.NewDecoder(r.Body).Decode(&fermentation); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if err := h.service.UpdateFermentation(&fermentation); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(fermentation)
// }

// func (h *FermentationHandler) DeleteFermentation(w http.ResponseWriter, r *http.Request) {
// 	// TODO
// }
