package main

import (
	"fmt"
	"errors"
	"reflect"
)

func main(){

	e1:=errors.New("EOF")
	e2:=errors.New("EOF")
	fmt.Println("e1==e2?:" ,e1==e2)
	v1:=reflect.ValueOf(e1)
	v2:=reflect.ValueOf(e2)
	fmt.Println("v1.Elem()==v2.Elem()",v1.Elem()==v2.Elem())

	fmt.Printf("e1 address %p\n",e1)
	fmt.Printf("e1 address %p\n",e2)
	fmt.Println("*e1==*e2?:" ,&e1==&e2)
}
