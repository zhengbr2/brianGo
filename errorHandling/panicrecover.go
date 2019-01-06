package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverName()
	if firstName == nil {
		panic(errors.New("runtime error: first name cannot be nil"))
	}
	if lastName == nil {
		fmt.Println("will trigger panic here")
		panic("runtime error: last name cannot be nil ")
	}
	panic(999)
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	lastName :="Mask"
	done := make(chan interface{})
	go func() {
		fullName(nil, &lastName)
		done <- 123
	}()

	<-done

	fullName(&firstName,nil)
	fullName(&firstName, &lastName)

	fmt.Println("returned normally from main")
}
