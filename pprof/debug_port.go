package main

import (
	"net/http"
	"runtime/pprof"
	"time"
)

func goppf() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":11181", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

var quit chan struct{} = make(chan struct{})

func f() {
	<-quit
}

func main() {
	for i := 0; i < 10000; i++ {
		go f()
	}

	go goppf() //启用跟踪查看
	for {
		time.Sleep(1 * time.Second)
	}
}
