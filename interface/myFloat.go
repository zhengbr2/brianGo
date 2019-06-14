package main

import (
	"fmt"
)

type ITest interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describeobj(t ITest) {
	fmt.Printf("Interface type %T value %v\n", t, t)
	t.Tester()

}

func main() {
	var t ITest
	f := MyFloat(89.7)
	t = &f
	describeobj( t)
	describeobj( f)
	//t.Tester()
}
