package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
	)

func main(){


	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	//log.Fatal(http.ListenAndServe(":8080",nil))

	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key",nil))

}


