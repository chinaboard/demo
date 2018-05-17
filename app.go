package main

import (
	"net/http"
	"strings"
)

type handle struct{}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	http.Handle("/", &handle{})
	http.ListenAndServe(":80", nil)
}
