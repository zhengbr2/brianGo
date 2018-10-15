package main

import (
	_ "fmt"
	"github.com/anacrolix/sync"
	"log"
	"net/rpc"
	"time"
	"runtime"
)

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

var (
	ThreadCount = 20
	Repeat      = 5000
)

func main() {

	runtime.GOMAXPROCS(4)
	before := time.Now()
	var wg sync.WaitGroup
	wg.Add(ThreadCount)
	for r := 0; r < ThreadCount; r++ {

		conn, err := rpc.Dial("tcp", "127.0.0.1:8095")
		if err != nil {
			log.Fatalln("dailing error: ", err)
		}
		go func() {
			for i := 0; i < Repeat; i++ {
				req := ArithRequest{9, 2}
				var res ArithResponse

				err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
				if err != nil {
					log.Fatalln("arith error: ", err)
				}
				//fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t := time.Now().Sub(before).Seconds()
	log.Println("total time:", t)
	log.Println("QPS:", float64(Repeat*ThreadCount)/t)

}
