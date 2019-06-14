package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var arr1 =[...]string{"hello","world","brian"}
	var arr3 =[...]string{"hello","world","brian","zheng"}
	var arr4 =[...]string{"hello","world","brian","zheng","go"}
	fmt.Println(arr1)
	arr2:=arr1
	fmt.Println(arr2)
	arr2[0]="Hi"
	string1 :="Hi"
	string2 :="Hi"
	fmt.Println("arr1:",arr1)
	fmt.Println("arr2:",arr2)
	fmt.Printf("string1 pointer %p:\n",&string1)
	fmt.Printf("string2 pointer %p:\n",&string2)
	fmt.Printf("arr1 pointer %p:",&arr1)
	fmt.Printf("arr2 pointer %p:",&arr2)
	fmt.Println("size of arr1:",unsafe.Sizeof(arr1))
	fmt.Println("size of arr2:",unsafe.Sizeof(arr2))
	fmt.Println("size of arr3:",unsafe.Sizeof(arr3))
	fmt.Println("size of arr4:",unsafe.Sizeof(arr4))
	fmt.Println("size of hi:",unsafe.Sizeof("hi"))
	fmt.Println("size of arr4[1]:",unsafe.Sizeof(arr4[1]))  //一个字符串长度 16bytes， 64位地址起点，64位长度，   16 ×5  = 80

}
