package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"role-based-to-do-api/internal/middleware"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.Handle("/protected", middleware.JWTAuthMiddleware(http.HandlerFunc(ProtectedEndpoint))).Methods("GET")

	r.HandleFunc("/sa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("as"))
	}).Methods("GET")

	s := r.PathPrefix("/lists").Subrouter()
	s.Use(middleware.JWTAuthMiddleware)

	s.HandleFunc("", CreateListHandler).Methods("POST")
	s.HandleFunc("", GetMyListsHandler).Methods("GET")

	return r
}
