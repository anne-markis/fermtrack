package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/anne-markis/fermtrack/internal/domain"
	"github.com/gorilla/mux"
)

type FermentationHandler struct {
	service *app.FermentationService
}

func NewFermentationHandler(service *app.FermentationService) *FermentationHandler {
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

func (h *FermentationHandler) CreateFermentation(w http.ResponseWriter, r *http.Request) {
	var fermentation domain.Fermentation
	if err := json.NewDecoder(r.Body).Decode(&fermentation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateFermentation(&fermentation); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FermentationHandler) UpdateFermentation(w http.ResponseWriter, r *http.Request) {
	// Similar to CreateFermentation
}

func (h *FermentationHandler) DeleteFermentation(w http.ResponseWriter, r *http.Request) {
	// Similar to GetFermentation
}
