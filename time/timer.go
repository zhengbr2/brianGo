package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	before:=time.Now()
	fmt.Println("now is:" , before)
	select {

	case res := <-c1:
		fmt.Println(res)
	case d:=<-time.After(time.Second * 1):
		fmt.Println("timeout 1")
		fmt.Println("duration is:" , d)
	}
	fmt.Println("ellapse duration:" , time.Now().Sub(before).Seconds())

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
