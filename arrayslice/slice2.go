package main

import (
	"fmt"
	"unsafe"
)

func main() {

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	pers := []string{}
	fmt.Println("sizeof empty slice:", unsafe.Sizeof(pers))  //24
	perA := [0]string{}
	fmt.Println("sizeof empty Array:", unsafe.Sizeof(perA)) //0
	fmt.Println("pers:", pers, "has new length", len(pers), "and capacity", cap(pers))
	pers = append(pers, "brian", "quan", "fey")
	fmt.Println("pers:", pers, "has new length", len(pers), "and capacity", cap(pers))
	fmt.Println("sizeof:", unsafe.Sizeof(pers))   //24

}
