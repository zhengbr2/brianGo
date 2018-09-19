package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "Foo"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("now sleep")
	time.Sleep(time.Second * 1)
	fmt.Println("now awake")

}
