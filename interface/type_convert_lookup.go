package main

import (
	"fmt"
)

func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func assert2(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	var s interface{} = 56
	assert(s)

	s = "Steven Paul"
	//assert(s) panic here, 转不成就挂掉了
	assert2(s)  //更稳妥

	findType("Naveen")
	findType(77)
	findType(89.98)
}
