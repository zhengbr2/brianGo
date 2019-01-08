package main

import (
	"fmt"
	"runtime"
	"github.com/pkg/errors"
)

func main() {
	call_1()
	outer_callerA()
	outer_callerB()

	e:=	errors.New("new Error")
	fmt.Printf("%+v",e)
}

func outer_callerA(){
	call_1()
}

func outer_callerB(){
	call_2()
}

func call_1() {
	var calldepth = 1;
	fmt.Println(runtime.Caller(calldepth))
}

func call_2() {
	var calldepth = 2;
	fmt.Println(runtime.Caller(calldepth))
}
