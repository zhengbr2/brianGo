package main

import "fmt"

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	p1 := person{name: "zzy", age: 100}
	p2 := person{name: "dj", age: 99}
	p3 := person{name: "px", age: 20}
	people := map[string]person{
		p1.name: p1,
		p2.name: p2,
		p3.name: p3,
	}
	//people["px"].isDead = true  not allowed， 本质还是：对值类型拷贝再修改其实无效
	//p3b:=people["px"]
	//p3b.isDead=true
	//fmt.Println(p3.isDead)

	whoIsDead(people)
	p3.isDead=true
	if p3.isDead {
		fmt.Println("who is dead?", p3.name)
	}

}

func whoIsDead(people map[string]person) {
	for name, _ := range people {
		if people[name].age < 50 {
			//people[name].isDead = true
			fmt.Println(people[name], people[name].isDead)

		}
	}
}
