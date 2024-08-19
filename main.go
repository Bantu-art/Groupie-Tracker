package main

import (
	"fmt"
	"groupie-tracker/routes"
	"net/http"
)

func main() {
	fmt.Println("Server running on :9000")

	http.HandleFunc("/", routes.Index)
	http.ListenAndServe(":9000", nil)
}
