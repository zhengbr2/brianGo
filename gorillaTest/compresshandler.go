package main

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

func main() {
	http.Handle("/gzip", useCompressHandler(handler2()))
	http.ListenAndServe(":1234", nil)
}

func handler2() http.Handler {
	return http.HandlerFunc(myHandler2)
}

func myHandler2(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}

func useCompressHandler(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}
