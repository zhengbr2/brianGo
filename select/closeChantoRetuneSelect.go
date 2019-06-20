package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	const n = 13

	finish := make(chan bool)
	var done sync.WaitGroup
	done.Add(n)

	for i := 0; i < n; i++ {

		go func() {
			defer done.Done()

			select {
			case <-time.After(1 * time.Minute):
			case <-finish:
				fmt.Println("return from selecct", time.Now())
			}

		}()
	}

	t0 := time.Now()

	time.Sleep(time.Second * 1)
	fmt.Println("will close now", time.Now())
	close(finish) // 关闭 finish 使其立即返回
	fmt.Println("closed", time.Now())
	done.Wait() // 等待所有的 goroutine 结束

	fmt.Printf("Waited %v for %d goroutines to stop\n", time.Since(t0), n)
	fmt.Println("end", time.Now())
}
