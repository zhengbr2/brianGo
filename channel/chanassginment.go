package main

import "fmt"
import "time"

func main() {
	inch := make(chan int)
	outch := make(chan int)
	go func() {
		var in <-chan int = inch
		var out chan<- int
		var val int
		for {
			select {
			case out <- val:
				println("step 2 coming from val:",val)
				out = nil // close channel out
				in = inch // open channel in
			case val = <-in:
				println("step 1 coming from inch:",val)
				out = outch //open channel out
				in = nil    //  close the channel in
			}
		}
	}()
	go func() {
		for r := range outch {
			fmt.Println("result:", r)
		}
	}()
	time.Sleep(0)
	inch <- 888
	inch <- 999
	time.Sleep(3 * time.Second)
}
