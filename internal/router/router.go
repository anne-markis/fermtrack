package router

import (
	"github.com/anne-markis/fermtrack/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(fermentationHandler *handlers.FermentationHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/fermentations", fermentationHandler.GetFermentations).Methods("GET")
	r.HandleFunc("/fermentations/{uuid}", fermentationHandler.GetFermentation).Methods("GET")
	r.HandleFunc("/fermentations", fermentationHandler.CreateFermentation).Methods("POST")

	return r
}
