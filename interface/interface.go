package main

import (
	"fmt"
	"time"
	"unsafe"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

var x interface{} = time.Now()

func main() {
	var phone Phone

	phone = new(NokiaPhone)

	phone.call()

	phone = new(IPhone)
	phone.call()

	fmt.Println("size of phone", unsafe.Sizeof(phone))
	fmt.Println("size of x", unsafe.Sizeof(x))
	var x interface{} = [3]int{1, 2, 3}
	fmt.Println("array is comparble, equal?:", x == x) //uncomparable for [] int

	var _ = (IntSet{}).String() // compile error: String requires *IntSet receiver
	var _ = (&IntSet{}).String() // compile error: String requires *IntSet receiver

	s := IntSet{}
	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method
}

type IntSet struct { /* ... */
}

func (*IntSet) String() string {
	return ""
}
