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


func main() {
	o := order{
		ordId:      456,
		customerId: 56,
		commodity:  "fruits",
	}

	showTypes()

	queryType(o)

	queryField(o)

	value:=reflect.ValueOf(o)

	fmt.Println("valueToInterface(value).commodity:",valueToInterface(value).commodity)

}

func showTypes(){

	ti := reflect.TypeOf(3)
	ts := reflect.TypeOf("string")
	tsli := reflect.TypeOf([]string{"string", "kk"})
	tinf := reflect.TypeOf(order{2, 3,"Fruits"})

	fmt.Println(ti, ts, tsli, tinf) //int, string, []string, main.order

}

func queryType( q interface{}){

	switch t2 := q.(type) {
	case order:
		fmt.Println("type is main.oder")
	case int:
		fmt.Println("type is int")
	default:
		fmt.Println(t2)

	}
}


func queryField(q interface{}) {

	println("reflect.TypeOf(q).Kind()== reflect.ValueOf(q).Kind():",reflect.TypeOf(q).Kind()== reflect.ValueOf(q).Kind())

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

func  valueToInterface(value reflect.Value ) order {

	return value.Interface().(order)

}