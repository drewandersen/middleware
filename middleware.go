package middleware

import (
	"log"
	"net/http"
)

func LogError(w http.ResponseWriter, err error, message string) {
	http.Error(w, "something went wrong", 500)
	log.Printf("error: %s: %v\n", message, err)
}

func HandleAccessLogs(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&sw, r)
		log.Printf("method: %s, proto: %s, path: %s, remote_addr: %s, status: %d\n", r.Method, r.URL.Path, r.Proto, r.RemoteAddr, sw.status)
	}
}

func AccessLogsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(&sw, r)
		log.Printf("method: %s, proto: %s, path: %s, remote_addr: %s, status: %d\n", r.Method, r.URL.Path, r.Proto, r.RemoteAddr, sw.status)
	})
}

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}
