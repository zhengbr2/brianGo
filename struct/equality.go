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
	fmt.Println("&a==&b?:", &a == &b)   //地址肯定不等啦

	var any interface{} = a
	var any2 interface{} = b
	fmt.Println("any==any2:", any == any2)
	//intreface的值是可比较的。如果两个interface 有相同的动态类型和相同的动态值，而且不为nil，那么这两个interface就是相同的

	fmt.Println("p1==p2:", &person{"brian", 30} ==&person{"brian", 30}) //false, memory address differ

	s1 := student{123, &person{"brian", 30}}
	s2 := student{123, &person{"brian", 30}}
	fmt.Println("s1.person==s2.person:", s1.person == s2.person)     // false, compara memory address
	fmt.Println("*s1.person==*s2.person:", *s1.person == *s2.person) // ture, compare value
	fmt.Println("s1==s2:", s1 == s2)                                 // false, since s1.person <> s2.person

	s3:=student{sid:123, person: &a}
	s4:=student{sid:123, person: &a}
	fmt.Println("s3==s4:", s3 == s4)                      //true, value equals
	test2()
}

func test2() {
	var user_id interface{}
	user_id = 123

	var id int
	id = 123

	//这里不能赋值，因为类型不一样  k
	//id = user_id
	id2 := user_id.(int)   //当然要转型啦

	//但是这里可以判断，为什么不同的类型可以判断相等？？？
	if user_id == id {
		fmt.Println("相等", user_id)
	} else {
		fmt.Println("不相等", user_id)
	}
	fmt.Println(id2==id)

	var flt_id interface{}
	flt_id=3.14
	if user_id == flt_id {
		fmt.Println("flt_id相等", user_id)
	} else {
		fmt.Println("flt_id不相等", user_id)
	}

	if 123 == 1234.0000{
		fmt.Println(" 123 == 123.33 equal")
	} else{
		fmt.Println(" 123 == 123.3 not equal")
	}

}
