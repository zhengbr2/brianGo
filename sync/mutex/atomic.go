package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0

	before := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {

			atomic.AddUint64(&ops, 1)

			runtime.Gosched()

		}()
	}

	//time.Sleep(time.Second * 5)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
	fmt.Println("cost:%f", time.Now().Sub(before).Seconds())
}
