package routes

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	data := indexTData{
		Title:   "Apis",
		Summary: "The home of artists, locations, dates and relations",
		Content: "Home page",
	}

	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
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
