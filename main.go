package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie-tracker/routes"
)

func main() {
	iport := port(8000)
	fmt.Printf("PORT: %d\n", iport)
	if len(os.Args) != 1 {
		log.Fatalf("Usage: go run .")
	}
	portStr := fmt.Sprintf(":%d", iport)
	fmt.Printf("Server listening on http://localhost%s\n", portStr)

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/static/", routes.Static)
	http.HandleFunc("/artists", routes.Artists)
	http.HandleFunc("/artists/", routes.Artist)
	http.HandleFunc("/dates/", routes.Dates)
	http.HandleFunc("/locations/", routes.Locations)
	http.HandleFunc("/relation/", routes.Relation)
	http.HandleFunc("/about/", routes.AboutHandler)
	http.HandleFunc("/contacts/", routes.Contacts)

	// error pages
	http.HandleFunc("/404", routes.Errors)
	http.HandleFunc("/405", routes.Errors)
	http.HandleFunc("/500", routes.Errors)

	http.ListenAndServe(portStr, nil)
}
