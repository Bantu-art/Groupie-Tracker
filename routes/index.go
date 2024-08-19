package routes

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	data := indexTData{
		Title:   "Apis",
		Summary: "The home of artists, locations, dates and relations",
		Content: "Home page",
	}

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	t := template.Must(templ, err)
	t.Execute(w, data)
}

type indexTData struct {
	Title   string
	Summary string
	Content any
}
