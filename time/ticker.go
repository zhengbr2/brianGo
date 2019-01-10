package main

import (
	"fmt"
	"time"
)

func main() {
	var chanTime <-chan time.Time
	ticker := time.NewTicker(time.Millisecond * 500)
	chanTime = ticker.C
	//chanTime = time.Tick(time.Millisecond * 500)

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

	time.Sleep(time.Millisecond * 4000)
	fmt.Println("Ticker stopped here ! ",time.Now())
	ticker.Stop()  // stop()之后， range结束
	time.Sleep(time.Millisecond * 1000)
}
