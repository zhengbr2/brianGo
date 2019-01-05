
package main

import (
"fmt"
"runtime"
"sync"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

func incCount2() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}
var mutex sync.Mutex

func incCount3() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}
