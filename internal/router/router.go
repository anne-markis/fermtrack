package router

import (
	"github.com/anne-markis/fermtrack/internal/handlers"
	"github.com/anne-markis/fermtrack/internal/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(
	fermentationHandler *handlers.FermentationHandler,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	// Public routes

	// login
	r.HandleFunc("/v1/login", authHandler.Login).Methods("POST")
	// new user
	r.HandleFunc("/v1/users", userHandler.CreateUser).Methods("POST")

	// Protected routes

	// user
	protectedRouter := r.PathPrefix("").Subrouter()
	protectedRouter.HandleFunc("/v1/users/{uuid}", userHandler.GetUser).Methods("GET")

	// fermentation operations
	protectedRouter.HandleFunc("/v1/fermentations/advice", fermentationHandler.GetFermentationAdvice).Methods("POST")
	protectedRouter.HandleFunc("/v1/fermentations", fermentationHandler.GetFermentations).Methods("GET")
	protectedRouter.HandleFunc("/v1/fermentations/{uuid}", fermentationHandler.GetFermentation).Methods("GET")

	// middleware
	r.Use(middleware.LoggingMiddleware) // all routes get logging
	protectedRouter.Use(middleware.AuthMiddleware)

	return r
}
