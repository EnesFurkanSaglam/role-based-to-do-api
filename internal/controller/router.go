package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/sa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("as"))
	}).Methods("GET")

	return r

}
