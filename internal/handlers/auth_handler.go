package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/anne-markis/fermtrack/internal/app"
	"github.com/rs/zerolog/log"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TODO test
type AuthHandler struct {
	authService *app.AuthService
}

func NewAuthHandler(authService *app.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	log.Error().Any("user", req.Username).Msg("successful login")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
