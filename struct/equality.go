package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	sid int
	*person
}

func main() {
	a := person{"brian", 30}
	b := person{"brian", 30}

	fmt.Println("a==b?:", a == b)
	fmt.Println("&a==&b?:", &a == &b)

	var any interface{}
	any = a
	var any2 interface{}
	any2 = a
	fmt.Println("any==any2:", any == any2)
	fmt.Println("any&==&any2:", &any == &any2)

	s1 := student{123, &person{"brian", 30}}
	s2 := student{123, &person{"brian", 30}}
	fmt.Println("s1.person==s2.person:", s1.person == s2.person)     // false, compara memory address
	fmt.Println("*s1.person==*s2.person:", *s1.person == *s2.person) // ture, compare value
	fmt.Println("s1==s2:", s1 == s2)                                 // false, since s1.person <> s2.person

}
