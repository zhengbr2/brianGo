package main

import (
	"fmt"
	"math/rand"
	"sync"
	"runtime"
)

var count int
var wg sync.WaitGroup

func main() {
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go read(i)
	}

	for i := 0; i < 5; i++ {
		go write(i)
	}

	wg.Wait()
}
//看不出有什么竞争冲突
func read(n int) {
	fmt.Printf("读goroutine %d 正在读取...\n", n)

	v := count
	runtime.Gosched()
	fmt.Printf("读goroutine %d 读取结束，值为：%d\n", n, v)
	wg.Done()
}

func write(n int) {
	fmt.Printf("写goroutine %d 正在写入...\n", n)
	v := rand.Intn(1000)

	count = v

	fmt.Printf("写goroutine %d 写入结束，新值为：%d\n", n, v)
	wg.Done()
}
