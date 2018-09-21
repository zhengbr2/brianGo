package main

import (
	"errors"
	"fmt"
)

func main() {
	//m =  map[string]string
	var m = make(map[string]string)
	m = map[string]string{}
	m["name"] = "zzy"

	var s = make([]int, 1)
	s[0] = 1
	//s[2] =2
	s = append(s, 2)

	var s2 []int
	s2 = append(s2, 1)
	fmt.Println(s2)

	fmt.Println(":=")
	i := 2
	if i > 1 {
		i, err := doDivision(i, 2)
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
	}
	fmt.Println(i) // no the same i

	fmt.Println("who is dead?")
	//
	//p1 := person{name: "zzy", age: 100}
	//p2 := person{name: "dj", age: 99}
	//p3 := person{name: "px", age: 20}
	//people := []person{p1, p2, p3}
	//whoIsDead(people)
	//for _, p := range people {
	//	if p.isDead {
	//		fmt.Println("who is dead?", p.name)
	//	}
	//}

	p1 := &person{name: "zzy", age: 100}
	p2 := &person{name: "dj", age: 99}
	p3 := &person{name: "px", age: 20}
	people := []*person{p1, p2, p3}
	whoIsDead2(people)
	for _, p := range people {
		if p.isDead {
			fmt.Println("who is dead?", p.name)
		}
	}

	p6 := person{name: "name", age: 100}
	p7 := p6
	p6.name = "changed"
	fmt.Println(p7.name)

}

func doDivision(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("input is invalid")
	}
	return x / y, nil
}

type person struct {
	name   string
	age    byte
	isDead bool
}

func whoIsDead(people []person) {
	for _, p := range people {
		if p.age < 50 {
			p.isDead = true
		}
	}
}

func whoIsDead2(people []*person) {
	for _, p := range people {
		if p.age < 50 {
			p.isDead = true
		}
	}
}
