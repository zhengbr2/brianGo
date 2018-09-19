package main

import "fmt"

import "unsafe"

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
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
}
