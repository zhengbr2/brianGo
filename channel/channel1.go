package main

import (
	"fmt"
)

func main() {

	done := make(chan bool, 1)
	go hello(done)
	//done<-true
	<-done
	fmt.Println("main function")
}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
