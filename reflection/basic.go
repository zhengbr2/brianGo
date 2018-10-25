package main

import (
	"fmt"
	"reflect"
)

type ifoo interface{
	 foo() string;
}

func foo() string{
	return "foo"
}

func main() {
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))


	fmt.Println("type: ", reflect.TypeOf(foo))
	fmt.Println("value: ", reflect.ValueOf(foo))


}


