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

	fmt.Println(time.Now())    // 0.0  时刻
	limiter := time.Tick(time.Millisecond * 200)  //0.2 时刻第一次打点
	time.Sleep(time.Second)
	fmt.Println(time.Now())  // 1.0 时刻
	for req := range requests {

		fmt.Println("request", req, <-limiter, "current time:", time.Now())
		// 1.0 时刻拿到 0.2
		// 1.2 时刻拿到 1.2
	}

	time.Sleep(time.Second)
	fmt.Println("outseide request", <-limiter, "current time:", time.Now())

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

		fmt.Println("request", i, <-burstyLimiter, "time:", time.Now())
	}
}
