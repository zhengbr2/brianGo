package main

import (
	"fmt"
)

func main() {

	//printChan()

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

func printChan() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T, value is %v\n", a, a)
	}
}
