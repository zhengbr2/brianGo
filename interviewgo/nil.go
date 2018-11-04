package main

import "fmt"

type IPeople interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live0() IPeople {

	return nil
}

func live1() IPeople {
	var stu *Student
	fmt.Printf("%h\n", &stu)
	fmt.Printf("%h\n", stu)
	// fmt.Printf("%h\n",*stu)  nil pointer here! panic
	return stu
}

func live2() IPeople {

	return &Student{}
}

type MyObject struct {
	name int
	count int
}
var ob MyObject

func main() {
	b0 := live0()
	_ = b0
	b1 := live1()
	b2 := live2()
	if b1 == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
	if b2 == nil {
		fmt.Println("AAAAAAA2")
	} else {
		fmt.Println("BBBBBBB2")
	}

	//fmt.Println("ob==nil?", ob==nil) error
}
