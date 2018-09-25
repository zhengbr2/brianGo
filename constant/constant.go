package main

import "fmt"

import (
	"unsafe"
	"log"
	"math"
)

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
	Zero1 =0
	Zero2 = 0.0
	T=3

)

const LENGTH int = 10
const WIDTH int = 5

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {

	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
	println(Unknown + Female + Male)

	
	log.Println("0.0==0?", Zero1==Zero2)
	log.Println( 3+0.1)
	math.Sin(Zero1)
	math.Sin(T)
	math.Sin(Zero2)
}
