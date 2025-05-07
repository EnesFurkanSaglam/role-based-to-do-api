package middleware

import (
	"context"
	"net/http"
	"role-based-to-do-api/internal/util"
	"strings"
)

type contextKey string

const (
	UserContextKey contextKey = "user"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {

			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokeStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := util.ParseJWT(tokeStr)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
