package main

import "fmt"

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	p := person{name: "zzy", age: 100}
	fmt.Println(p)
	isDead(&p)
	fmt.Println(p)
}

func isDead(p interface{}) {

	if p.(*person).age < 101 { //解指针
		p.(*person).isDead = true
	}
}
