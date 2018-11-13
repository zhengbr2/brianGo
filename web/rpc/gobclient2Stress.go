package main

import (
	"flag"
	"github.com/anacrolix/sync"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"

	"math/rand"
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
	ThreadCount = 5000
	Repeat      = 100
)

func init() {
	println("Usage Sample:")
	println(os.Args[0] + " -thread 100  -repeat 5000")
	flag.IntVar(&ThreadCount, "thread", 1000, "how many threads(goroutine) running in client side")
	flag.IntVar(&Repeat, "repeat", 500, "repeat count within one thread")
	flag.Parse()
	println("your input: -thread:" + strconv.Itoa(ThreadCount) + " -repeat:" + strconv.Itoa(Repeat))
}

func main() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	_ = rd

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
				req := ArithRequest{ i/2, i}
				var res ArithResponse

				err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
				if err != nil {
					log.Fatalln("arith error: ", err)
				}
				//time.Sleep( time.Millisecond * time.Duration(rd.Intn(1000)));
				//
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
