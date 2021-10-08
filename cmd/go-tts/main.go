package main

import (
	"fmt"
	"net/http"
)

func main() {
	service := New()

	http.HandleFunc("/create", service.createHandler)
	http.HandleFunc("/play", service.playHandler)
	fmt.Println("Server listening on port 8090")
	http.ListenAndServe(":8090", nil)
}
