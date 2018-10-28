package webutils

import (
	"log"
	"net/http"
)

func LogRequest(r http.Request) {
	log.Printf("type: request, method: %s, url: %s, remote_addr: %s", r.Method, r.URL, r.RemoteAddr)
}

func LogError(w http.ResponseWriter, err error, message string) {
	http.Error(w, "something went wrong", 500)
	log.Printf("error: %s: %v\n", message, err)
}
