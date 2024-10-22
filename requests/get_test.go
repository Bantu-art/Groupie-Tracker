package requests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.String() {
		case "/success":
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "Mock response body")
		case "/error":
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, "Internal Server Error")
		default:
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Not Found")
		}
	}))
	defer mockServer.Close()

	tests := []struct {
		url             string
		expectedSuccess bool
	}{
		{
			url:             mockServer.URL + "/success",
			expectedSuccess: true,
		},
		{
			url:             mockServer.URL + "/error",
			expectedSuccess: true,
		},
		{
			url:             "http://invalid-url",
			expectedSuccess: false,
		},
	}

	for _, test := range tests {
		t.Run(test.url, func(t *testing.T) {
			// ignore the incoming data
			// we have a guarantee that if remote server does not return an error, then data will be useful
			// all we have to account for is the status code and the boolean that are returned by the Get function
			_, _, success := Get(test.url)

			if success != test.expectedSuccess {
				t.Errorf("expected success %v, got %v", test.expectedSuccess, success)
			}
		})
	}
}
