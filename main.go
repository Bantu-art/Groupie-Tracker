package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/routes"
)

func main() {
	fmt.Println("Server running on :9000")

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/artists", routes.Artists)
	http.HandleFunc("/artists/", routes.Artist)
	http.ListenAndServe(":9000", nil)
}
