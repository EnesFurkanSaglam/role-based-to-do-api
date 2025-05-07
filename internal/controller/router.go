package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"role-based-to-do-api/internal/middleware"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/sa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("as"))
	}).Methods("GET")

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.Handle("/protected", middleware.JWTAuthMiddleware(http.HandlerFunc(ProtectedEndpoint))).Methods("GET")

	s := r.PathPrefix("/lists").Subrouter()
	s.Use(middleware.JWTAuthMiddleware)

	s.HandleFunc("", CreateListHandler).Methods("POST")
	s.HandleFunc("", GetMyListsHandler).Methods("GET")
	s.HandleFunc("/steps", AddStepHandler).Methods("POST")
	s.HandleFunc("/steps", GetStepsHandler).Methods("GET")
	s.HandleFunc("/steps/complete", CompleteStepHandler).Methods("POST")
	s.HandleFunc("/steps/delete", DeleteStepHandler).Methods("POST")
	s.HandleFunc("/delete", DeleteListHandler).Methods("POST")
	s.HandleFunc("/lists/update", UpdateListHandler).Methods("POST")
	s.HandleFunc("/steps/update", UpdateStepHandler).Methods("POST")

	return r
}
