package main

import (
	"context"
	"fmt"
	"time"
)

type q struct{}

func main() {

	ctx, cancel := context.WithCancel(context.Background())   // context.Background().Deadline() 0,false
	done := make(chan struct{})
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				done <- struct{}{}
				return
			default:
				fmt.Println("goroutine监控中...")
				tm, b:=ctx.Deadline()
				fmt.Println("goroutine监控中..." , tm,b)
				time.Sleep(2 * time.Second)
			}
		}

		//close(done)
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	tm, b:=ctx.Deadline()
	fmt.Println("goroutine监控中2..." , tm,b)
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	//time.Sleep(5 * time.Second)
	<-done

}
