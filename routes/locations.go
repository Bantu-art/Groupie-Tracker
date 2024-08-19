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

func Locations(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/locations/")
	if id == "" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	locationUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
	var location ctypes.Locations

	jsonLocationData, status, found := requests.Get(locationUrl)
	if found && status == "200 OK" {
		err := json.Unmarshal([]byte(jsonLocationData), &location)
		if err != nil {
			fmt.Println("Error unmarshalling location JSON:", err)
			return
		}

		data := struct {
			Title     string
			Locations []string
		}{
			Title:     "Locations",
			Locations: location.Locations,
		}

		templ, err := template.ParseFiles("templates/locations.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		t := template.Must(templ, err)
		t.Execute(w, data)
	} else {
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}
