package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	wg := sync.WaitGroup{}
	wg.Add(50)
	for i := 0; i < 40; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
		//time.Sleep(1)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
