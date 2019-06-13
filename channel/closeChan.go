package main

import (
	"fmt"
	"time"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl) // must close it if no more msg write into chan
	fmt.Println("Closed")
}
func main() {
	ch := make(chan int)
	go producer(ch)

	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("break now! ", v, ok)
			break
		}
		fmt.Println("Received ", v, ok)
	}


	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Received ", v)
	}

	chanint := make(chan int)
	close(chanint)
	select {
	case i := <-chanint:
		println("read i:", i)
	}
	select {
	case i := <-chanint:
		println("read i:", i)
	}

	chanint = nil
	select {
	case i := <-chanint:
		println("read i:", i)
	default:
		println("never read i!")
	}

}
