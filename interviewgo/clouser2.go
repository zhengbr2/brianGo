package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 40; i++ {   //会导致输出的i不可控
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()

	}
	time.Sleep(100)
	fmt.Println("-----------------------------------------")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("in clouser i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
