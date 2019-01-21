package main

import (
	"reflect"
	"unsafe"
	"fmt"
)

func main(){

	str := "helloWorld"
	str2 :=str[:]   //str2 and str3 是一样的。 为了减少字符串的复制，如果没有改动， 就沿用旧的底层byte数组
	str3 :=str

	strheader:=(*reflect.StringHeader)(unsafe.Pointer(&str))
	strhearder2:=(*reflect.StringHeader)(unsafe.Pointer(&str2))
	strhearder3:=(*reflect.StringHeader)(unsafe.Pointer(&str3))



	fmt.Printf("%#v\n",strheader)
	fmt.Printf("%#v\n",strhearder2)
	fmt.Printf("%#v\n",strhearder3)

	sli := []byte(str)
	bsliheader:=(*reflect.SliceHeader)(unsafe.Pointer(&sli))   //changed
	fmt.Printf("%#v\n",bsliheader)
}

