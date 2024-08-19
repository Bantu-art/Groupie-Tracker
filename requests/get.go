package requests

import (
	"fmt"
	"io"
	"net/http"
)

// Function to perform a GET request
func Get(url string) (data string, status string, success bool) {
	resp, err := http.Get(url)
	if err != nil {
		data = fmt.Sprintf("Error on fetching resource: %q", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		data = fmt.Sprintf("Error reading response body: %q", err.Error())
		return
	}

	status = resp.Status
	data = string(body)
	success = true
	return
}
