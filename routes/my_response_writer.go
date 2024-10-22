package routes

import (
	"bytes"
	"net/http"
)

/*
* create a custom ResponseWriter  to capture the response for testing
* It needs to implement thehttp.ResponseWriter interface
* Since a correct status code implies the body was rendered correctly, there is no need to check for the contents of the rendered html
 */
type MyResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Body       *bytes.Buffer
}

func (w *MyResponseWriter) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *MyResponseWriter) Write(b []byte) (int, error) {
	// we are not testing the contents of the body right now
	// w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}
