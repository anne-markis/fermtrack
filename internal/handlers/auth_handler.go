package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/anne-markis/fermtrack/internal/app/domain"
	"github.com/rs/zerolog/log"
)

// TODO test
type AuthHandler struct {
	authService *app.AuthService
	userRepo    domain.UserRepository
}

func NewAuthHandler(authService *app.AuthService, userRepo domain.UserRepository) *AuthHandler {
	return &AuthHandler{authService: authService, userRepo: userRepo}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Any("err", err.Error()).Msg("invalid request")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		log.Error().Any("err", err.Error()).Msg("invalid request")
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		log.Error().Any("err", err.Error()).Msg("failed login")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	log.Error().Any("username", req.Username).Any("user_uuid", user.UUID).Msg("successful login")

	w.Header().Set("Content-Type", "application/json")

	response := LoginResponse{
		Token:    token,
		Username: req.Username,
		UUID:     user.UUID,
	}
	json.NewEncoder(w).Encode(response)
}
