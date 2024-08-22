package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/requests"
	"groupie-tracker/types"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}

	var data artistsTData
	jsonData, status, found := requests.Get("https://groupietrackers.herokuapp.com/api/artists")
	if found && status == "200 OK" {
		var artists []types.Artist
		err := json.Unmarshal([]byte(jsonData), &artists)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}

		data = artistsTData{
			Title:   "Apis | Artists",
			Summary: "The home of artists",
			Content: artists,
		}
	} else {
		data = artistsTData{
			Title:   "Artists",
			Summary: "The home of artists",
		}
	}

	templ, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}
	t := template.Must(templ, err)
	t.Execute(w, data)
}

type artistsTData struct {
	Title   string
	Summary string
	Content any
}
