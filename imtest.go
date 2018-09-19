package main

import (
	"runtime"
	"fmt"
)

func main(){
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(8)
}
