package middleware

import (
	configs "4-order-api/config"
	"4-order-api/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/initiate" || r.URL.Path == "/auth/verify" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if token == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userPhone", data.Phone)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
