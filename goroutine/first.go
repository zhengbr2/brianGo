package main

import (
	"fmt"
)

func hello() {
	//time.Sleep(time.Second*1)
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	fmt.Println("main function")
}
