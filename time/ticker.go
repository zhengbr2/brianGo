package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	//time.Sleep(time.Millisecond * 4000)
	//ticker.Stop()
	<-make(chan int)

	fmt.Println("Ticker stopped")
}
