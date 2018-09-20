package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(1 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}


func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine!")
	go hello(done)
	<-done
	fmt.Println("Main received data")

	a := make(chan int, 2)
	a <- 1
	a <- 1
	fmt.Println("after write a")
	inta := <-a
	inta = <-a

	_ = inta
	fmt.Println("after read a")

	i := make(chan int)
	ci := make(chan chan int, 10)
	ci <- i
	ci <- i
	ci <- i
	ri := <-i
	fmt.Println("after read i", ri)
	fmt.Println("after read ci")

}
