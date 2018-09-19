package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	//mean = float32(sum)/count    // no auto convert
	fmt.Printf("mean 的值为: %f\n", mean)

	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小

}
