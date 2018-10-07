package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"runtime/debug"
)

var (
	host = flag.String("host", "", "host name or IP")
	port = flag.String("port", "3333", "port #")
)

func main() {
	var (
		err      error
		checkErr = func(err error) {
			if err != nil {
				debug.PrintStack()
				os.Exit(1)
			}
		}
	)

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
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}
