package main

import (
	"fmt"
	"time"
)

func main() {
	var chanTime <-chan time.Time
	ticker := time.NewTicker(time.Second * 1)
	chanTime = ticker.C
	chanTime = time.Tick(time.Millisecond * 500)

	go func() {
		for t := range chanTime {
			fmt.Println("Tick at 1", t)
		}
	}()

	go func() {
		for t := range chanTime {
			fmt.Println("Tick at 2", t)
		}
	}()

	//go func() {
	//	time.Sleep(time.Second * 3)
	//	chanTime=time.Tick(time.Millisecond * 500)  // not work...
	//
	//}()

	//time.Sleep(time.Millisecond * 4000)
	//ticker.Stop()
	<-make(chan int)

	fmt.Println("Ticker stopped")
}
