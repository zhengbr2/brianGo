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

	var any3 interface{}
	any3 = b
	fmt.Println("any==any2:", any == any2)
	fmt.Println("any&==&any2:", &any == &any2)
	//intreface的值是可比较的。如果两个interface 有相同的动态类型和相同的动态值，而且不为nil，那么这两个interface就是相同的
	fmt.Println("any&==&any3:", &any == &any3)

	s1 := student{123, &person{"brian", 30}}
	s2 := student{123, &person{"brian", 30}}
	fmt.Println("s1.person==s2.person:", s1.person == s2.person)     // false, compara memory address
	fmt.Println("*s1.person==*s2.person:", *s1.person == *s2.person) // ture, compare value
	fmt.Println("s1==s2:", s1 == s2)                                 // false, since s1.person <> s2.person

	test2()
}

func test2() {
	var user_id interface{}
	user_id = 123

	var id int
	id = 123

	//这里不能赋值，因为类型不一样
	//id = user_id

	//但是这里可以判断，为什么不同的类型可以判断相等？？？
	if user_id == id {
		fmt.Println("相等", user_id)
	} else {
		fmt.Println("不相等", user_id)
	}
}
