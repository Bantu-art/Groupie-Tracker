package routes

import (
	"html/template"
	"net/http"
)

// AboutHandler handles the request for the about page
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}

	// Try to parse the about page template
	aboutTmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		// About template is missing, redirect to 500 error page
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// If about template exists, render it normally
	data := struct {
		Title   string
		Summary string
	}{
		Title:   "About Groupie Tracker",
		Summary: "Learn about our service",
	}

	if err := aboutTmpl.Execute(w, data); err != nil {
		// If template execution fails, redirect to 500 error page
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
}
