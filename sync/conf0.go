package main

import (
	"fmt"
	sync "sync"
	"time"
)

var locker = new(sync.Mutex)
var cond2 = sync.NewCond(locker)

func test(x int) {
	cond2.L.Lock() //获取锁
	fmt.Println("aaa: ", x)
	cond2.Wait() //等待通知  暂时阻塞
	fmt.Println("bbb: ", x, " time:", time.Now())
	//time.Sleep(time.Second * 1)
	cond2.L.Unlock() //释放锁
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond2.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond2.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond2.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 3)
	fmt.Println("finish all")

}
