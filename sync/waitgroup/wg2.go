package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Millisecond)

	wg.Done()
	wg.Wait() //wg between neibought threads //brian I don't think anybody will write codes like this.
	fmt.Printf("Goroutine %d ended\n", i)
}

func main() {
	no := 30000
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	fmt.Println("wg.wait() here")
	wg.Done()
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
