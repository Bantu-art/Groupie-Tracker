package routes

import (
	"net/http"
	"os"
)

/*
* Static: A function that provides a layer over the default go servefile function to prevent directory listings
*/
func Static(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if info.IsDir() {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
