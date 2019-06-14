package main

import (
	"time"
	"fmt"
)

func main() {
	test := map[int]int {1:1}

	go func() {
		i := 0
		for i < 10000 {
			test[1]=1
			i++
		}
	}()

	go func() {
		i := 0
		for i < 10000 {
			test[1]=1
			i++
		}
	}()

	time.Sleep(2*time.Second)
	fmt.Println(test)
}
