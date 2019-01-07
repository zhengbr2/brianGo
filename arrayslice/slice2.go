package main

import (
	"fmt"
	"unsafe"
	"reflect"
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

	n1 := make ([] int, 0,5)
	n1 = append(n1, 1,2)
	n2 := n1
	n2=append(n2, 3)
	fmt.Println(n1)   // [1,2]
	fmt.Println(n2)   // [1,2,3]
	fmt.Printf("address of n1:%p\n",n1)
	fmt.Printf("address of n2:%p\n",n2)

	sh1:=(*reflect.SliceHeader)(unsafe.Pointer(&n1))
	fmt.Printf("A1 Data:%d,Len:%d,Cap:%d\n",sh1.Data,sh1.Len,sh1.Cap)

	sh2 :=(*reflect.SliceHeader)(unsafe.Pointer(&n2))
	fmt.Printf("A Data:%d,Len:%d,Cap:%d\n", sh2.Data, sh2.Len, sh2.Cap)

	n3:=[]int {1,2}
	fmt.Println("n3",n3)
	n3=append(n3, 3)
	fmt.Println("n3",n3)

	m1:= map[int]string {
		1:"brian",
		2:"quan",
	}

	m1[3]="Ken"
	fmt.Println(m1)



}
