package main

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	addr := "localhost:8972"

	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")

	timeClient := client.NewXClient("TimeS", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer timeClient.Close()
	for i := 0; i < 100000; i++ {

		tm := time.Now().Add(time.Hour * 3)
		log.Println("pricessing request" + strconv.Itoa(i))

		var ret time.Time
		errT := timeClient.Call(context.Background(), "Time", tm, &ret)
		if errT != nil {
			log.Println("encountered error:", errT)
		}
		log.Println("result from Time:", ret)
	}
}
