package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}
func main() {
	sparrow := &Bird{"Sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem() // Elem() for pointer.. if not pointer, no need Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v | %v \n", i, typeOfT.Field(i).Name, f.Type(), f, f.Interface())
		//f.Interface() turn to interface again
	}

	for i := 0; i < typeOfT.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, typeOfT.Field(i).Type, f.Interface())
	}

	{
		sparrow := Bird{"Sparrow", 3}
		s := reflect.ValueOf(sparrow) // Elem() for pointer.. if not pointer, no need Elem()
		typeOfT := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f) //why no need Interface()

		}
	}
}
