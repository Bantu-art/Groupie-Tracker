package routes

import (
	"html/template"
	"net/http"
)

// AboutHandler handles the request for the about page
func Contacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}

	// Try to parse the contacts page template
	contactsTmpl, err := template.ParseFiles("templates/contacts.html")
	if err != nil {
		// Contacts template is missing, redirect to 500 error page
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// If contacts template exists, render it normally
	data := struct {
		Title   string
		Summary string
	}{
		Title:   "Contacts",
		Summary: "Get in touch with us",
	}

	if err := contactsTmpl.Execute(w, data); err != nil {
		// If template execution fails, redirect to 500 error page
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}
