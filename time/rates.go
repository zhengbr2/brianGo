package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	fmt.Println(time.Now())
	limiter := time.Tick(time.Millisecond * 200)
	time.Sleep(time.Second)
	for req := range requests {

		fmt.Println("request", req, <-limiter,  "current time:" ,time.Now())
	}

	time.Sleep(time.Second)
	fmt.Println("outseide request",  <-limiter, "current time:" ,time.Now())

	fmt.Println("----------------")
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()


	for i := 1; i <= 5; i++ {

		fmt.Println("request", i, <-burstyLimiter, "time:",time.Now())
	}
}
