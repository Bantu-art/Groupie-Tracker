package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAboutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// call the AboutHandler with the Response Recorder and the request we just built
	handler := http.HandlerFunc(AboutHandler)
	handler.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	}
}
