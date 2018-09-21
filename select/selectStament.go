package main

import (
	"fmt"
	"time"
)

func main() {
	var c1, c2, c3 = make( chan int), make( chan int),make( chan int)
	var i1, i2 int = 11, 22

	go  func(){
		c1 <- 100
	}()

	go  func(){
		fmt.Println("received c2", <-c2)
	}()

	go  func(){
		c3 <- 100
		//close(c3)
	}()

	//time.Sleep(10* time.Millisecond)
	select {
	case i1 = <-c1:
		fmt.Println("case1")
		fmt.Println("received ", i1, " from c1")
	case c2 <- 88:
		fmt.Println("case2")
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		fmt.Println("case3")
		if ok {
			fmt.Println("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}


	}

	time.Sleep(10* time.Millisecond)
}
