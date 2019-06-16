package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
	"github.com/anacrolix/sync"
)

func main() {

	var ops uint64 = 0
	var wg sync.WaitGroup

	wg.Add(1000000)
	before := time.Now()
	for i := 0; i < 1000000; i++ {
		go func() {

			atomic.AddUint64(&ops, 1)
			//ops++

			runtime.Gosched()
			wg.Done()

		}()
	}


	wg.Wait()
	opsFinal := atomic.LoadUint64(&ops)   // ops: 1,000,000
	fmt.Println("ops:", opsFinal)
	fmt.Printf("cost:%f", time.Now().Sub(before).Seconds())
}
