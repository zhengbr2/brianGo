package main

import (
	"brianGo/web/grpc/example"
	"context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"

	"strconv"
)

// 定义请求地址
const (
	ADDRESS string = "localhost:8080"
)

var wg sync.WaitGroup

// main 方法实现对 gRPC 接口的请求
var now=time.Now()
func main() {


	wg.Add(100)
	for i := 0; i < 100; i++ {
		go dial(i)
	}
	wg.Wait()

}
func dial(j int) {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + ADDRESS)
	}
	for i := 0; i < 1000; i++ {
		client := example.NewFormatDataClient(conn)
		resp, err := client.DoFormat(context.Background(), &example.Data{Text: "hello,world:"+ strconv.Itoa(j*1000+i)})
		if err != nil {
			log.Fatalln("Do Format error:" + err.Error())
		}
		log.Println(resp.Text)
	}
	conn.Close()
	wg.Done()
	pass:=time.Now().Sub(now).Seconds()
	log.Println("elapse:",pass)
}
