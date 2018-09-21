package main

import "fmt"

type person struct {
	name   string
	age    int
	isDead bool
}

func main() {
	p1 := &person{name: "zzy", age: 100}
	p2 := &person{name: "dj", age: 99}
	p3 := &person{name: "px", age: 20}
	people := map[string]*person{
		p1.name: p1,
		p2.name: p2,
		p3.name: p3,
	}
	people["px"].isDead = true

	whoIsDead3(people)
	if p3.isDead {
		fmt.Println("who is dead?", p3.name)
	}
}

func whoIsDead3(people map[string]*person) {
	for name, _ := range people {
		if people[name].isDead {
			people[name].age = -1
			fmt.Println(people[name], people[name].isDead)

		}
	}
}
