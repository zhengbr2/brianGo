package main

import (
	"fmt"
	"unsafe"
)

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	fmt.Println(unsafe.Sizeof(t)) //0
	t.ShowA()
	t.ShowB()
}
