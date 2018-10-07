package main

import (
	"fmt"
	"net"
	"log"
	"bytes"
	"encoding/binary"
)

//const (
//	BYTES_SIZE uint16 = 1024
//	HEAD_SIZE  int    = 2
//)

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		log.Println("accepting a new connection")
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		log.Println("dispatching a go routine to handle the connection")
		go doConn(conn)
	}
}

func doConn(conn net.Conn) {
	var (
		buffer           = bytes.NewBuffer(make([]byte, 0, BYTES_SIZE))
		bytes            = make([]byte, BYTES_SIZE);
		isHead      bool = true
		contentSize int
		head        = make([]byte, HEAD_SIZE)
		content     = make([]byte, BYTES_SIZE)
	)
	for {
		log.Println("entering first infinite for loop")
		readLen, err := conn.Read(bytes);
		if err != nil {
			log.Println("Error reading", err.Error())
			return
		}
		_, err = buffer.Write(bytes[0:readLen])
		if err != nil {
			log.Println("Error writing to buffer", err.Error())
			return
		}

		for {
			log.Println("entering second infinite for loop")
			if isHead {
				if buffer.Len() >= HEAD_SIZE {
					_, err := buffer.Read(head)
					if err != nil {
						fmt.Println("Error reading", err.Error())
						return
					}
					contentSize = int(binary.BigEndian.Uint16(head))
					isHead = false
				} else {
					log.Println("breaking second infinte loop in A")
					break
				}
			}
			if !isHead {
				if buffer.Len() >= contentSize {
					_, err := buffer.Read(content[:contentSize])
					if err != nil {
						fmt.Println("Error reading", err.Error())
						return
					}
					fmt.Println(string(content[:contentSize]))
					isHead = true
				} else {
					log.Println("breaking second infinte loop in B")
					break
				}
			}
		}
	}
}


//func main(){
//	StartServer("localhost:50002")
//}