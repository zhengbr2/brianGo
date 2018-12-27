package main

import "fmt"

func main(){
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("i is :", i,  " x is : ",x) // "0" "2" "4" "6" "8"
		case ch <- i:

		}
	}
}

