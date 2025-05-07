package main

import (
	"fmt"
	"net/http"
	"role-based-to-do-api/internal/controller"
)

func main() {
	r := controller.NewRouter()
	fmt.Println("8082 port")
	err := http.ListenAndServe(":8082", r)
	if err != nil {
		panic(err)
	}
}
