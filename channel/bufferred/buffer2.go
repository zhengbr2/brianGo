package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch", time.Now())
	}
	close(ch)
	fmt.Println("channel is closed here ", time.Now())
}
func main() {
	ch := make(chan int, 2) // if buffer is 1, then read out 4 numbers, instead of 5
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch", time.Now())
		time.Sleep(1 * time.Second)

	}
}
