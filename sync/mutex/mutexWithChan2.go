package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
}
type writeOp struct {
	key int
	val int
}

var b = make(chan int)

func main() {

	var ops int64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				b <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				b <- 0
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				read := &readOp{
					key: rand.Intn(5),
				}
				reads <- read
				total += <-b
				fmt.Println(total)
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
				}
				writes <- write
				<-b
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
