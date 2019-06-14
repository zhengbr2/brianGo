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
var x2 interface{} = nil

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	interfaceCompare()
}

func interfaceCompare(){
	var phone Phone
	phone = new(NokiaPhone)

	fmt.Println("size of phone", unsafe.Sizeof(phone))   //16   8bytes for value , and 8 bytes for type
	fmt.Println("size of NokiaPhone", unsafe.Sizeof(NokiaPhone{}))   //0
	fmt.Println("size of x", unsafe.Sizeof(x)) //16
	fmt.Println("size of x2", unsafe.Sizeof(x2)) //16

	var x interface{} = [3]int{1, 2, 3}
	fmt.Println("array is comparble, equal?:", x == x) //uncomparable for [] int

	 //_ = (IntSet{}).String() // compile error: String requires *IntSet receiver
	ss:=IntSet{}
	_=ss.String()   // okay
	 _ = (&IntSet{}).String() // okay

	s := IntSet{}
	var sg fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s  // compile error: IntSet lacks String method
	fmt.Println("size of Stringer", unsafe.Sizeof(sg))   //16   8bytes for value , and 8 bytes for type
}

type IntSet struct { /* ... */
}

func (*IntSet) String() string {
	return ""
}
