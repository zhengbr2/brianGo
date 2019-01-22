package main

import (
	"brianGo/web/rpcx/raw"
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

func main() {

	log.SetFlags(log.Lmicroseconds)
	addr := "localhost:8972"
	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")

	arithCall(d)

	echoCall(d)

	timeQuery(d)
}

func arithCall(d client.ServiceDiscovery) {

	args := raw.Args{5, 6}
	reply := &raw.Reply{}
	arithClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer arithClient.Close()
	log.Printf("args:%v", args)
	err := arithClient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Println("encountered error:", err)
	}
	log.Println("result for Mul:", reply.C)
	err = arithClient.Call(context.Background(), "Add", args, reply)
	if err != nil {
		log.Println("encountered error:", err)
	}
	log.Println("result for Add:", reply.C)
}

func echoCall(d client.ServiceDiscovery) {

	echoClient := client.NewXClient("Echo", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer echoClient.Close()
	args2 := "brian"
	log.Printf("args:%v", args2)
	var reply2 string
	err2 := echoClient.Call(context.Background(), "Echo", args2, &reply2)
	if err2 != nil {
		log.Println("encountered error:", err2)
	}
	log.Println("result from Echo:", reply2)
}

func timeQuery(d client.ServiceDiscovery) {

	timeClient := client.NewXClient("TimeS", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer timeClient.Close()

	tm := time.Now().Add(time.Hour * 3)
	log.Println(tm)

	var ret time.Time

	errT := timeClient.Call(context.Background(), "Time", tm, &ret)
	if errT != nil {
		log.Println("encountered error:", errT)
	}
	log.Println("result from Time:", ret)

}
