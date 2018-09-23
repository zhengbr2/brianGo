package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
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
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer)
	println(f,c)
}
