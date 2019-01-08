package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()

	default:
		fmt.Printf("unknown type %v \n", i)
	}
}

func main() {
	findType("Naveen")

	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)

	var w io.Writer
	w = os.Stdout
	f,ok := w.(*os.File) // success: f == os.Stdout
	fmt.Println("\n", f, ok)
	c,ok := w.(*bytes.Buffer)
	fmt.Println("\n", c,ok)
}
