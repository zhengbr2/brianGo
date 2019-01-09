package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//t := time.Now().UnixNano()
	NumPublishers := 4 //runtime.NumCPU()
	totalIterations := int64(1000 * 1000 * 20)
	iterations := totalIterations / int64(NumPublishers)

	channel := make(chan int64, 1024*64)
	var wg sync.WaitGroup
	wg.Add(NumPublishers + 1)
	var readerWG sync.WaitGroup
	readerWG.Add(1)
	for i := 0; i < NumPublishers; i++ {
		go func() {
			fmt.Println(time.Now(), "goroutine",i)
			wg.Done() // no use?
			wg.Wait() //
			fmt.Println(time.Now(), "goroutine after wait",i)
			for i := int64(0); i < iterations; {
				select {
				case channel <- i:
					i++
				default:
					continue
				}
			}

		}()
	}
	go func() {
		for i := int64(0); i < totalIterations; i++ {
			select {
			case  <-channel:


			default:
				continue
			}
		}
		readerWG.Done()
	}()
	//time.Sleep(time.Second * 1 )
	wg.Done()
	t := time.Now().UnixNano()
	wg.Wait()
	readerWG.Wait()
	t = (time.Now().UnixNano() - t) / 1000000 //ms
	fmt.Printf("opsPerSecond: %d\n", totalIterations*1000/t)
	fmt.Println(time.Now(),"time elapse: %d\n", t)
}
