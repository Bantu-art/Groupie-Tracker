package routes

import (
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request) {
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
