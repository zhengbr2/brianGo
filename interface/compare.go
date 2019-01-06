package main

import (
	"errors"
	"fmt"
	"syscall"
)

func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func main() {

	e1 := errors.New("EOF")
	e2 := errors.New("EOF")
	fmt.Println("e1==e2?:", e1 == e2)  //false!
	fmt.Printf("e1 address %p\n", e1)
	fmt.Printf("e1 address %p\n", e2)

	var i1 interface{} = "EOF"
	var i2 interface{} = "EOF"
	fmt.Println("i1==i2?:", i1 == i2)  //true


	e3 := New("shit")
	e4 := New("shit")
	e3o := e3.(*errorString)
	e4o := e4.(*errorString)
	fmt.Println("e3.(errorString)==e4.(errorString)", e3o == e4o) // pointer, won't equal

	e5 := struct{ name string }{"brian"}
	e6 := struct{ name string }{"brian"}
	fmt.Println("e5==e6", e5 == e6)

	e7 := &struct{ name string }{"brian"}
	e8 := &struct{ name string }{"brian"}
	fmt.Println("e7==e8", e7 == e8)

}

func errorNum() {

	er1a := syscall.Errno(1)
	er1b := syscall.Errno(1)

	fmt.Println(er1a) //"Incorrect function"
	fmt.Println("er1a==er1b", er1a == er1b)
	fmt.Println("interface:", er1a) //"Incorrect function"
}
