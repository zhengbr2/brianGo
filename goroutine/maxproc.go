package main

import (
	"fmt"
	"runtime"
	"sync"
	"log"
)

func main() {
	log.Println("number of cpu",runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		defer wg.Done()
		for i:=1;i<1000;i++ {
			fmt.Println("A:",i)

		}
	}()
	go func(){
		defer wg.Done()
		for i:=1;i<1000;i++ {
			fmt.Println("B:",i)

		}
	}()
	wg.Wait()
}