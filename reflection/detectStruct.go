package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
	commodity  string
}

func createQuery(q interface{}) {

	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		t := reflect.TypeOf(q)

		fmt.Println("kind same:", v.Kind() == t.Kind())
		fmt.Println("Number of fields", v.NumField())     // NumField() only for struct type: otherwise panic: reflect: NumField of non-struct type
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d, name:%s, type:%s, value:%v\n", i, t.Field(i).Name, t.Field(i).Type, v.Field(i))

		}
	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
		commodity:  "fruits",
	}
	createQuery(o)
	println("reflect.TypeOf(o).Kind()== reflect.ValueOf(o).Kind():",reflect.TypeOf(o).Kind()== reflect.ValueOf(o).Kind())

}
