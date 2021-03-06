package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"brianGo/web/rpcx2/example"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	s := server.Server{}
	s.RegisterName("Arith", new(example.Arith), "")
	go s.Serve("tcp", *addr)
	select {}
}
