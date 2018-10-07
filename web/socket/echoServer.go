package main

import (
	"flag"
	"fmt"
	_ "io"
	"net"
	"os"
	"runtime/debug"
	"io"
)

var (
	host = flag.String("host", "", "host name or IP")
	port = flag.String("port", "3333", "port #")

	err      error
	checkErr = func(err error) {
		if err != nil {
			fmt.Println(err)
			debug.PrintStack()
			os.Exit(1)
		}
	}
)

func main() {

	flag.Parse()
	var l net.Listener
	l, err = net.Listen("tcp", *host+":"+*port)
	checkErr(err)

	defer l.Close()
	fmt.Println("Listening on " + *host + ":" + *port)
	for {
		conn, err := l.Accept()
		checkErr(err)

		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	//var buf = make([]byte,1024)
	defer conn.Close()

		for {
			fmt.Printf("expected only entring here for 1 time \n")
			count, _ := io.Copy(conn, conn)
			fmt.Printf("received chars:%d\n", count)
		}
		//n,_:=conn.Read(buf)
		//conn.Write(buf[0:n])

}
