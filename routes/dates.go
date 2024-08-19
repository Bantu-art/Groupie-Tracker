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
    t := template.Must(template.ParseFiles("templates/dates.html"))
    
    id := strings.TrimPrefix(r.URL.Path, "/dates/")
    if id == "" {
        http.Error(w, "Page not found", http.StatusNotFound)
        return
    }
    
    dateUrl := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
    var date ctypes.Dates
    
    jsonDateData, status, found := requests.Get(dateUrl)
    if found && status == "200 OK" {
        err := json.Unmarshal([]byte(jsonDateData), &date)
        if err != nil {
            fmt.Println("Error unmarshalling location JSON:", err)
            return
        }
        
        data := struct{
			Title string
			Dates []string
		}{
            Title: "Dates",
            Dates: date.Dates,
        }
        
        t.Execute(w, data)
    } else {
        http.Error(w, "Page not found", http.StatusNotFound)
    }
}
