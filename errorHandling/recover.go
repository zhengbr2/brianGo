package main

import (
	"fmt"
		"time"
	"runtime/debug"
)

func r() {
	if r := recover(); r != nil {
		fmt.Printf("%T:\n", r)
		fmt.Println("Recovered an panic:", r)
		fmt.Println("debug output call stack")
		time.Sleep(time.Millisecond*2)
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
	defer func() {
		fmt.Println("defer func() expected to executed before the end return")
	}()

	fmt.Println("will trigger an panic here")
	time.Sleep(time.Millisecond*2)
	go a()

	time.Sleep(time.Millisecond * 100)
	fmt.Println("normally returned from main")
}
