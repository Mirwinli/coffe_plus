package core_http_response

import "net/http"

var (
	statusCodeUninitialized = -1
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		statusCode:     statusCodeUninitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode

	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *ResponseWriter) GetStatusCode() int {
	return rw.statusCode
}
