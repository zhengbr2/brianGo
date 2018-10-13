package main

import (
	"github.com/smallnest/rpcx/client"
	"context"
	"brianGo/web/rpcx/raw"
	"log"
	"time"
)


func main(){
	addr:="localhost:8972"
	args:=raw.Args{5,6}
	reply:=&raw.Reply{}

	d := client.NewPeer2PeerDiscovery("tcp@"+addr, "")

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	log.Printf("args:%v",args)
	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err!= nil{
		log.Println("encountered error:",err)
	}
	log.Println("result for Mul:",reply.C)

	err=xclient.Call(context.Background(), "Add", args, reply)
	if err!= nil{
		log.Println("encountered error:",err)
	}
	log.Println("result for Add:",reply.C)



	echo := client.NewXClient("Echo", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer echo.Close()

	args2:="brian"
	log.Printf("args:%v",args2)
	var reply2 string

	err2 := echo.Call(context.Background(), "Echo", args2, &reply2)
	if err2!= nil{
		log.Println("encountered error:",err2)
	}
	log.Println("result from Echo:",reply2)


	{

		tServer := client.NewXClient("TimeS", client.Failtry, client.RandomSelect, d, client.DefaultOption)
		defer tServer.Close()

		tm:=time.Now().Add(time.Hour * 3)
		log.Println(tm)

		var ret time.Time

		errT := tServer.Call(context.Background(), "Time", tm, &ret)
		if errT != nil{
			log.Println("encountered error:", errT)
		}
		log.Println("result from Time:",ret)
	}
}
