package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var b bool = false
	fmt.Println(b)
	var i int32 = 0
	fmt.Println(unsafe.Sizeof(b))  //1
	fmt.Println(unsafe.Sizeof(i))  //4
	var ch byte = 65

	fmt.Println( string(rune(ch)))
	var r rune = 'C'
	fmt.Println( r)  //rune is int32
	fmt.Println( string(r))
	fmt.Println(unsafe.Sizeof(ch))   //unit8
	fmt.Println(unsafe.Sizeof(r))
}


