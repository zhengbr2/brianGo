package main

import (
	"reflect"
	"fmt"
)

type order struct {
	ordId      int
	customerId int
	commodity string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
		commodity: "fruits",
	}
	createQuery(o)
}