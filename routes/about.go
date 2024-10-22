package routes

import (
	"html/template"
	"net/http"
)

// AboutHandler handles the request for the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	// Define the path to your about page template
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// You can pass any data to the template if needed
	data := struct {
		Title string
	}{
		Title: "About Groupie Tracker",
	}

	// Render the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
