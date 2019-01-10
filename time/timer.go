package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "interrupted 1"
	}()

	before := time.Now()
	fmt.Println("starting time is:", before.Format("2006-01-02 15:04:05"))
	chanTM := time.After(time.Second * 1)
	go func(tm <-chan time.Time) {
		fmt.Println("read chanTM again:", <-tm)
	}(chanTM)

	select {

	case res := <-c1:
		fmt.Println(res)
	case tm := <-chanTM:
		fmt.Println("timeout 1")
		fmt.Println("current time is:", tm.Format("2006-01-02 15:04:05"))
	}
	fmt.Println("ellapse duration:", time.Now().Sub(before).Seconds())

}
