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
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/locations/")
	if id == "" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	locationUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
	var location ctypes.Locations

	jsonLocationData, status, found := requests.Get(locationUrl)
	if found && status == "200 OK" {
		err := json.Unmarshal([]byte(jsonLocationData), &location)
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusFound)
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
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}
		t := template.Must(templ, err)
		t.Execute(w, data)
	} else {
		http.Redirect(w, r, "/500", http.StatusFound)
	}
}
