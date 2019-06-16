package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)

}

func main() {
	no := 3
	var wg sync.WaitGroup
	wg.Add(no)
	for i := 0; i < no; i++ {
		go process(i, &wg)
	}
	fmt.Println("wg.wait() here")
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
