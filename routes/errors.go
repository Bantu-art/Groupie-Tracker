package routes

import (
	"html/template"
	"net/http"
)

func Errors(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	data := struct {
		Code    int
		Message string
	}{
		Code:    500,
		Message: "There was an error on our side",
	}
	switch url {
	case "/404":
		{
			data.Code = http.StatusNotFound
			data.Message = "Page not found"
		}
	case "/500":
		{
			data.Code = http.StatusInternalServerError
			data.Message = "Internal server error"
		}
	case "/405":
		{
			data.Code = http.StatusMethodNotAllowed
			data.Message = "Method not allowed"
		}
	}

	w.WriteHeader(data.Code)
	t := template.Must(template.ParseFiles("templates/errors.html"))
	t.Execute(w, data)
}
