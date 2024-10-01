package router

import (
	"github.com/anne-markis/fermtrack/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(fermentationHandler *handlers.FermentationHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/fermentations/advice", fermentationHandler.GetFermentationAdvice).Methods("POST")
	r.HandleFunc("/v1/fermentations", fermentationHandler.GetFermentations).Methods("GET")
	r.HandleFunc("/v1/fermentations/{uuid}", fermentationHandler.GetFermentation).Methods("GET")

	return r
}
