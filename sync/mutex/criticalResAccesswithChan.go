package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}
func main() {

	before := time.Now()

	var w sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()

	fmt.Println("final value of x", x)
	fmt.Println("cost:%f", time.Now().Sub(before).Seconds())

}
