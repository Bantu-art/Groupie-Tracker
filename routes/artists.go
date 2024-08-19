package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/types"
	"groupie-tracker/requests"
)

func Artists(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/artists.html"))

	var data artistsTData
	jsonData, status, found := requests.Get("https://groupietrackers.herokuapp.com/api/artists")
	if found && status == "200 OK" {
		var artists []types.Artist
		err := json.Unmarshal([]byte(jsonData), &artists)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
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
	t.Execute(w, data)
}

type artistsTData struct {
	Title   string
	Summary string
	Content any
}