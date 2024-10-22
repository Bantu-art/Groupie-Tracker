package routes

import (
	"html/template"
	"net/http"
)

// AboutHandler handles the request for the about page
func Contacts(w http.ResponseWriter, r *http.Request) {
	// Define the path to your contact page template
	tmpl, err := template.ParseFiles("templates/contacts.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// You can pass any data to the template if needed
	data := struct {
		Title string
	}{
		Title: "Contact Groupie Tracker",
	}

	// Render the template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
