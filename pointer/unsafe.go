package main

import (
	"fmt"
	"unsafe"
)

func main() {

	test()

	u := new(user)

	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))   // both convert to uintptr before add
	*pAge = 20

	fmt.Println(*u)

}

type user struct {
	name string
	age  int
}

func test() {

	i := 9
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))  // any pointer can be converted to unsafe.Pointer, and vice verses

	*fp = *fp * 3

	fmt.Println(i)
}
