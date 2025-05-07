package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("8082 port")

	http.HandleFunc("/sa", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("as"))
	})
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
