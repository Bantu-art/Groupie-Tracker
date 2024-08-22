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
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}
	filePath := "." + r.URL.Path
	info, err := os.Stat(filePath)
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}

	if info.IsDir() {
		http.Redirect(w, r, "/404", http.StatusFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
