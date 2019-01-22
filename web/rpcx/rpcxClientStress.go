package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"sync"
	"time"
)

var (
	CLIENT_NUM   = 4
	SINGLE_COUNT = 100000
)

func main() {
	log.SetFlags(log.Lmicroseconds)

	addr := "localhost:8972"
	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")
	before := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(CLIENT_NUM)

	for i := 0; i < CLIENT_NUM; i++ {
		go queryTime(d, &wg)
	}

	wg.Wait()
	log.Println("consumed time %f:", time.Now().Sub(before).Seconds())
	log.Println("average %f:", float64(CLIENT_NUM*SINGLE_COUNT)/time.Now().Sub(before).Seconds())
}

func queryTime(d client.ServiceDiscovery, wg *sync.WaitGroup) {

	timeClient := client.NewXClient("TimeS", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	defer func() { timeClient.Close(); wg.Done() }()

	for i := 0; i < SINGLE_COUNT; i++ {

		tm := time.Now().Add(time.Hour * 3)
		var ret time.Time
		errT := timeClient.Call(context.Background(), "Time", tm, &ret)
		if errT != nil {
			log.Println("encountered error:", errT)
		}
	}
}
