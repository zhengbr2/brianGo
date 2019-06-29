package main

import (
	"context"
	"fmt"
	"time"
)

type q struct{}

func main() {

	fmt.Println(time.Now())
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	go func(ctx context.Context) {
		ctxsub, _ := context.WithCancel(ctx)

		for {
			select {
			case <-ctxsub.Done():
				fmt.Println("监控退出，停止了...",time.Now())
				done <- struct{}{}
				return
			default:
				fmt.Println("goroutine监控中...",time.Now())
				time.Sleep(2 * time.Second)
			}
		}


		//close(done)
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("可以了，通知监控停止",time.Now())
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)
	<-done

}
