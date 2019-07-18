package main

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/pprof"
	"time"
)

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func main() {

	StasticGroutine := func() {
		for {
			time.Sleep(1e9)
			total := runtime.NumGoroutine()
			fmt.Println("current goroutine count:", total)
		}
	}

	go StasticGroutine()

	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	for i := 0; i < 100; i++ {
		doWork(nil)
		fmt.Println("do work:", i)
	}
	// 这里还有其他任务执行
	http.HandleFunc("/dump", dumpHandler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Done.")
}
