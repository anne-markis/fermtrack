package middleware

import (
	"net/http"
	"strings"

	"github.com/anne-markis/fermtrack/internal/utils"
	"github.com/rs/zerolog/log"
)

// AuthMiddleware Enforces a valid JWT token in the Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Error().Msg("missing authorization header")
			http.Error(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := utils.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			log.Error().Any("err", err).Msg("invalid token")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
