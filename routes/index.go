package routes

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))

	data := indexTData{
		Title:   "Apis",
		Summary: "The home of artists, locations, dates and relations",
		Content: "Home page",
	}
	t.Execute(w, data)
}

type indexTData struct {
	Title   string
	Summary string
	Content any
}
