package controller

import (
	"encoding/json"
	"net/http"
	"role-based-to-do-api/internal/middleware"
	"role-based-to-do-api/internal/service"
	"role-based-to-do-api/internal/util"
)

type CreateListRequest struct {
	Name string `json:"name"`
}

func CreateListHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*util.Claims)

	var req CreateListRequest
	json.NewDecoder(r.Body).Decode(&req)

	list := service.CreateList(req.Name, claims.Username)
	json.NewEncoder(w).Encode(list)
}

func GetMyListsHandler(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*util.Claims)

	lists := service.GetUserLists(claims.Username, claims.Role)
	json.NewEncoder(w).Encode(lists)
}
