package main

import (
	"net/http"
	"io"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	http.Handle("/",useLoggingHandler(http.HandlerFunc(myHandler)))
	http.ListenAndServe(":1234",nil)
}


func myHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw,"Hello World")
}


func useLoggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout,next)
}
