package main

import "fmt"

func main() {
	p:=person{name:"张三"}
	fmt.Printf("address outsite modify:%p\n", &p)
	p.modify()  //复制传值
	fmt.Println(p.String()) //复制传值

}

type person struct {
	name string
}


func (p person) String() string{
	fmt.Printf("address in String:%p\n", &p)
	return "the person name is "+p.name
}

func (p *person) modify(){
	fmt.Printf("address in modify:%p\n", &p)
	p.name = "李四"
}
