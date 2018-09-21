package main

type person struct {
	name   string
	age    byte
	isDead bool
}

func main() {
	p := person{name: "zzy", age: 100}
	isDead(&p)
}

func isDead(p interface{}) {
	if p.(*person).age < 101 {
		p.(person).isDead = true
	}
}
