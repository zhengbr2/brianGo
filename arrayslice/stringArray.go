package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var arr1 =[...]string{"hello","world","brian"}
	fmt.Println(arr1)
	arr2:=arr1
	fmt.Println(arr2)
	arr2[0]="Hi"
	fmt.Println("arr1:",arr1)
	fmt.Println("arr2:",arr2)
	fmt.Println("size of arr1:",unsafe.Sizeof(arr1))
	fmt.Println("size of arr2:",unsafe.Sizeof(arr2))


}
