package main

import (
	"fmt"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered an panic", r)
		debug.PrintStack()
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a, but will not be executed!")

}

func main() {
	fmt.Println("entring into main funciton")
	defer func(){
		fmt.Println("defer func() expected to executed before the end return")
	}()

	fmt.Println("will trigger an panic here")
	a()
	fmt.Println("normally returned from main")
}
