package main

import (
	"reflect"
	"fmt"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	fmt.Printf("Type %T\n", q)
	fmt.Printf("Type %t\n", q)


	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)


	ti:= reflect.TypeOf(3)
	ts:= reflect.TypeOf("string")
	tsli:= reflect.TypeOf([]string {"string","kk"})
	tinf:= reflect.TypeOf(order{2,3})

	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println(ti,ts, tsli, tinf)

	switch t2 :=q.(type){
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
