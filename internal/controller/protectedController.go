package controller

import (
	"fmt"
	"net/http"
	"role-based-to-do-api/internal/middleware"
	"role-based-to-do-api/internal/util"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserContextKey)
	if user == nil {
		http.Error(w, "User info not take", http.StatusInternalServerError)
		return
	}

	claims := user.(*util.Claims)
	msg := fmt.Sprintf("Hiiii my mann  %s youre role :%s ", claims.Username, claims.Role)
	w.Write([]byte(msg))
}
