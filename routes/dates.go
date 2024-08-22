package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"groupie-tracker/requests"
	ctypes "groupie-tracker/types"
)

func Dates(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/dates/")
	if id == "" {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	dateUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	var date ctypes.Dates

	jsonDateData, status, found := requests.Get(dateUrl)
	if found && status == "200 OK" {
		err := json.Unmarshal([]byte(jsonDateData), &date)
		if err != nil {
			fmt.Println("Error unmarshalling location JSON:", err)
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}

		data := struct {
			Title string
			Dates []string
		}{
			Title: "Dates",
			Dates: date.Dates,
		}

		templ, err := template.ParseFiles("templates/dates.html")
		if err != nil {
			http.Redirect(w, r, "/500", http.StatusFound)
			return
		}
		t := template.Must(templ, err)
		t.Execute(w, data)
	} else {
		http.Redirect(w, r, "/500", http.StatusFound)
        return
	}
}
