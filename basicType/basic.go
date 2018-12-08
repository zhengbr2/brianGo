package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var b bool = false
	fmt.Println(b)
	var i int32 = 0
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(i))
	var ch byte = 65

	println( string(rune(ch)))
	var r rune = 'C'
	println( r)
	println( string(r))
	fmt.Println(unsafe.Sizeof(ch))
	fmt.Println(unsafe.Sizeof(r))
}


