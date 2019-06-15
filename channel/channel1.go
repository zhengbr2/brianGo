package main

import (
	"fmt"
	"unsafe"
)

func main() {

	done := make(chan bool, 1)
	go hello(done)
	//done<-true
	<-done
	fmt.Println("main function")
	fmt.Printf("size of chanel:%d",unsafe.Sizeof(done))
}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}
