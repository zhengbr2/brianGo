package main

import "fmt"

type ipeople interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() ipeople {
	var stu *Student
	return stu
}

func main() {
	b := live()
	if b == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
