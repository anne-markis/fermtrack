package router

import (
	"github.com/anne-markis/fermtrack/internal/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(
	fermentationHandler *handlers.FermentationHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// auth
	r.HandleFunc("/v1/login", authHandler.Login).Methods("POST")

	// user
	r.HandleFunc("/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/v1/users/{uuid}", userHandler.GetUser).Methods("GET")

	// fermentation operations
	r.HandleFunc("/v1/fermentations/advice", fermentationHandler.GetFermentationAdvice).Methods("POST")
	r.HandleFunc("/v1/fermentations", fermentationHandler.GetFermentations).Methods("GET")
	r.HandleFunc("/v1/fermentations/{uuid}", fermentationHandler.GetFermentation).Methods("GET")

	return r
}
