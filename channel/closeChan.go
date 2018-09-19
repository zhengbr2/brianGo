package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
	fmt.Println("Closed")
}
func main() {
	ch := make(chan int)
	go producer(ch)
	time.Sleep(time.Second * 1)
	fmt.Println("read after close??? no... must read before close")
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		time.Sleep(time.Second * 1)
		fmt.Println("Received ", v)
	}
}
