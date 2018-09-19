package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	finish := make(chan struct{})

	var done sync.WaitGroup

	done.Add(1)

	go func() {

		fmt.Println("enter select now", time.Now())
		select {

		case <-time.After(1 * time.Hour):

		case <-finish:

		}


		done.Done()
		fmt.Println("done now", time.Now())

	}()

	time.Sleep(time.Second)
	t0 := time.Now()

	close(finish)

	done.Wait()
	fmt.Println("will close now", time.Now())
	fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))

}
