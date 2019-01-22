package main

import (
	service "brianGo/web/rpcx2/example"
	"context"
	"flag"
	"github.com/anacrolix/sync"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

var (
	addr        = flag.String("addr", "127.0.0.1:8972", "server address")
	ThreadCount = 100
	RepeatCount = 1000
)

func main() {
	Peer2Peer()
}
func Peer2Peer() {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	before := time.Now()

	var wg sync.WaitGroup
	wg.Add(ThreadCount)
	for r := 0; r < ThreadCount; r++ {

		go func() {
			xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
			defer xclient.Close()
			for i := 0; i < RepeatCount; i++ {
				args := &service.Args{
					A: 10,
					B: 20,
				}

				reply := &service.Reply{}
				err := xclient.Call(context.Background(), "Mul", args, reply)
				if err != nil {
					log.Fatalf("failed to call: %v", err)
				}

				//log.Printf("%d * %d = %d", args.A, args.B, reply.C)

			}
			wg.Done()
		}()
	}
	wg.Wait()
	t := time.Now().Sub(before).Seconds()
	log.Println("total time:", t)
	log.Println("QPS:", float64(RepeatCount*ThreadCount)/t)
}
