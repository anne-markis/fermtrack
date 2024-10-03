package middleware

import (
	"net/http"
	"strings"

	"github.com/anne-markis/fermtrack/internal/utils"
	"github.com/rs/zerolog/log"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/v1/login" || (r.RequestURI == "/v1/users" && r.Method == http.MethodPost) { // TODO dislike, unsafe, shameful
			log.Info().Msg("allowing auth bypass")
			next.ServeHTTP(w, r)
			return
		}
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
