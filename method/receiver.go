package main

import "fmt"

type person struct {
	name string
}

func (p person) String() string{
	return "the person name is "+p.name
}

func main() {
	p:=person{name:"Original"}
	(p).modify1() //值接收者，修改无效
	fmt.Println(p)
	(&p).modify1() //值接收者，修改无效
	fmt.Println(p)
	p.modify2()  //值接收者，修改成功
	fmt.Println(p)
	(&p).modify2() //值接收者，修改成功
	fmt.Println(p)
}


func (p person) modify1(){
	p.name = "updated"
}

func (p *person) modify2(){
	p.name = "updated"
}