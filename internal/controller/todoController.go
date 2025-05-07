package controller

import (
	"encoding/json"
	"fmt"
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

type CreateStepRequest struct {
	ListID  int    `json:"list_id"`
	Content string `json:"content"`
}

func AddStepHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateStepRequest
	json.NewDecoder(r.Body).Decode(&req)

	step := service.AddStep(req.ListID, req.Content)
	json.NewEncoder(w).Encode(step)
}

type IDRequest struct {
	ID int `json:"id"`
}

func CompleteStepHandler(w http.ResponseWriter, r *http.Request) {
	var req IDRequest
	json.NewDecoder(r.Body).Decode(&req)

	ok := service.CompleteStep(req.ID)
	if !ok {
		http.Error(w, "Step not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Completed"))
}

func DeleteStepHandler(w http.ResponseWriter, r *http.Request) {
	var req IDRequest
	json.NewDecoder(r.Body).Decode(&req)

	ok := service.DeleteStep(req.ID)
	if !ok {
		http.Error(w, "Step not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Deleted"))
}

func GetStepsHandler(w http.ResponseWriter, r *http.Request) {
	listID := r.URL.Query().Get("list_id")
	if listID == "" {
		http.Error(w, "list_id not found", http.StatusBadRequest)
		return
	}

	var id int
	fmt.Sscanf(listID, "%d", &id)

	steps := service.GetSteps(id)
	json.NewEncoder(w).Encode(steps)
}

func DeleteListHandler(w http.ResponseWriter, r *http.Request) {
	var req IDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ok := service.DeleteList(req.ID)
	if !ok {
		http.Error(w, "List not found or already deleted", http.StatusNotFound)
		return
	}

	w.Write([]byte("List deleted (soft delete)."))
}
