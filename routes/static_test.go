package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestStatic(t *testing.T) {
	testFile := "test_exists.txt" // create a file for testing => it will exist when tests are checking the directory
	err := os.WriteFile(testFile, []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)

	tests := []struct {
		method     string
		url        string
		statusCode int
	}{
		{"GET", "/test_exists.txt", http.StatusOK}, // (file exists) AND (valid HTTP method) => should be served
		{"POST", "/test_exists.txt", http.StatusFound}, // (file exists) BUT (wrong HTTP method) => do not return files when method is wrong
		{"GET", "/test_notexists_file.txt", http.StatusFound}, // (file does not exist) AND (valid HTTP method)
		{"GET", "/static/", http.StatusFound}, // (directory) AND (valid HTTP method) => test if directory is being served
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.url, nil)
		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(Static)

		handler.ServeHTTP(recorder, req)

		if recorder.Code != test.statusCode {
			t.Errorf("expected status %d, got %d", test.statusCode, recorder.Code)
		}
	}
}
