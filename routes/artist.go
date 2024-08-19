package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	ctypes "groupie-tracker/types"
	 "groupie-tracker/requests"
)

func Artist(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	if len(urlParts) != 3 {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	id := urlParts[len(urlParts)-1]
	artistUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)

	t := template.Must(template.ParseFiles("templates/artist.html"))

	// template data type
	var data artistTData

	// request status
	var status string
	var found bool

	// fetch artist
	var jsonArtistData, jsonRelationData string
	jsonArtistData, status, found = requests.Get(artistUrl)
	if found && status == "200 OK" {
		// found artist
		var artist ctypes.Artist
		err := json.Unmarshal([]byte(jsonArtistData), &artist)
		if err != nil {
			fmt.Println("Error unmarshalling artist JSON:", err)
			return
		}

		// fetch relation (date => location)
		relationUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)
		var relation ctypes.Relation
		jsonRelationData, status, found = requests.Get(relationUrl)
		if found && status == "200 OK" {
			// found locations
			err := json.Unmarshal([]byte(jsonRelationData), &relation)
			if err != nil {
				fmt.Println("Error unmarshalling location JSON:", err)
				return
			}
		}

		data = artistTData{
			Title:    "An artists tale",
			Summary:  "A detailed view of an artist",
			Artist:   artist,
			Relation: relation,
		}
	} else {
		data = artistTData{
			Title:   "Apis | Artists",
			Summary: "The home of artists",
		}
	}

	t.Execute(w, data)
}

type artistTData struct {
	Title    string
	Summary  string
	Artist   ctypes.Artist
	Relation ctypes.Relation
}
