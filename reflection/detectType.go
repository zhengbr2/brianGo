package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {

	fmt.Printf("Type %+#v\n", q)

	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	k2 := v.Kind()
	println(k2 == k) //true

	fmt.Println("Type ", t) // main.order
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k) //struct

	ti := reflect.TypeOf(3)
	ts := reflect.TypeOf("string")
	tsli := reflect.TypeOf([]string{"string", "kk"})
	tinf := reflect.TypeOf(order{2, 3})

	fmt.Println(ti, ts, tsli, tinf) //int string []string main.order

	switch t2 := q.(type) {
	case order:
		fmt.Println("type is oder")
	case int:
		fmt.Println("type is int")
	default:
		fmt.Println(t2)

	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}
