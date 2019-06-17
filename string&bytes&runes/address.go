package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

func main(){

	str := "hello,世界"
	str2 :=str[:]   //str2 and str3 是一样的。 为了减少字符串的复制，如果没有改动， 就沿用旧的底层byte数组
	str3 :=str[6:]
	sli4 :=[]byte(str)
	sli5 :=[]byte(str3)

	strheader:=(*reflect.StringHeader)(unsafe.Pointer(&str))
	strhearder2:=(*reflect.StringHeader)(unsafe.Pointer(&str2))
	strhearder3:=(*reflect.StringHeader)(unsafe.Pointer(&str3))
	strhearder4:=(*reflect.StringHeader)(unsafe.Pointer(&sli4))
	strhearder5:=(*reflect.StringHeader)(unsafe.Pointer(&sli5))


	//以下结果都是一样
	fmt.Printf("%#v\n",strheader)
	fmt.Printf("%#v\n",strhearder2)

	fmt.Printf("%#v\n",strhearder3)

	//以下结果都不一样
	fmt.Printf("%#v\n",strhearder4)
	fmt.Printf("%#v\n",strhearder5)

	fmt.Printf("size %d\n", unsafe.Sizeof(str))   //16
	fmt.Printf("size %d\n", unsafe.Sizeof(str3))  //16
	fmt.Printf("size %d\n", unsafe.Sizeof(sli5))  //24  slice多了一个capacity...因为str不可变 len == cap ..故省略一个8byte
	// str3[1]='E' now allow assigne to.

	sli := []byte(str)
	bsliheader:=(*reflect.SliceHeader)(unsafe.Pointer(&sli))   //changed
	fmt.Printf("%#v\n",bsliheader)
}

