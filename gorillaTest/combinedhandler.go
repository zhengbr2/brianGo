package main

import (
	"net/http"
	"io"
	"os"
	"github.com/gorilla/handlers"
)



func main() {
	http.Handle("/",useCombinedLoggingHandler(handler3()))
	http.ListenAndServe(":1234",nil)
}


func handler3() http.Handler{
	return http.HandlerFunc(myHandler3)
}

func myHandler3(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw,"Hello World")
}


func useCombinedLoggingHandler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout,next)
}
