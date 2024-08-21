package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/routes"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatalf("Usage: go run .")
	}
	fmt.Println("Server running on http://localhost:9000/")
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/static/", routes.Static)
	http.HandleFunc("/artists", routes.Artists)
	http.HandleFunc("/artists/", routes.Artist)
	http.HandleFunc("/dates/", routes.Dates)
	http.HandleFunc("/locations/", routes.Locations)
	http.HandleFunc("/relation/", routes.Relation)

	http.ListenAndServe(":9000", nil)
}
