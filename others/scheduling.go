package main

import (
	"fmt"
	"runtime"
)

func main() {
	done := false
	go func() {
		done = true
	}()
	for !done {

		//fmt.Println("not done!")
		runtime.Gosched()
	}
	fmt.Println("done!")

}
