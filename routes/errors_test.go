package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		url          string
		expectedCode int
	}{
		{"/404", http.StatusNotFound},
		{"/500", http.StatusInternalServerError},
		{"/405", http.StatusMethodNotAllowed},
		{"/unknown", 500},
	}

	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.url, nil)
			recorder := httptest.NewRecorder()
			w := &MyResponseWriter{
				ResponseWriter: recorder,
				StatusCode:     http.StatusOK,
				Body:          new(bytes.Buffer),
			}

			// since it uses templates/errors.html to switch between error pages
			// ensure that templates/errors.html
			Errors(w, req)

			res := recorder.Result()
			if res.StatusCode != test.expectedCode {
				t.Errorf("Expected status code to be %d, but got %d", test.expectedCode, res.StatusCode)
			}
		})
	}
}
