package main

import (
	"context"
	"fmt"
	"time"
)

var key1 string = "key1"
var key2 string = "key2"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key1, "value1")
	valueCtx = context.WithValue(valueCtx, key2, "value2")

	go watch(valueCtx, key1)
	go watch(valueCtx, key2)
	time.Sleep(7 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(2 * time.Second)
}

func watch(ctx context.Context, key string) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}
