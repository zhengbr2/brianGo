package main

import (
	"net/http"
	"io"
	"os"
	"github.com/gorilla/handlers"
)



func main() {
	http.Handle("/",useLoggingHandler(handler()))
	http.ListenAndServe(":1234",nil)
}

func handler() http.Handler{
	return http.HandlerFunc(myHandler)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw,"Hello World")
}


func useLoggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout,next)
}
