package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)

	{
		var data *byte
		var in interface{}
		fmt.Println(data, data == nil) //prints: <nil> true
		fmt.Println(in, in == nil)     //prints: <nil> true
		in = data    // interface has type and value, here type is not nil. but value is nil...
		fmt.Println(in, in == nil) //prints: <nil> false
		//'data' is 'nil', but 'in' is not 'nil'
	}

}
