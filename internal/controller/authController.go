package controller

import (
	"encoding/json"
	"net/http"
	"role-based-to-do-api/internal/repository"
	"role-based-to-do-api/internal/util"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user, ok := repository.FindUser(req.Username, req.Password)
	if !ok {
		http.Error(w, "username or password wrong", http.StatusUnauthorized)
		return
	}

	token, err := util.GenerateJWT(user.Username, user.Password)
	if err != nil {
		http.Error(w, "token generation error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
