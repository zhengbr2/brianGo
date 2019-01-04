package main

import "fmt"

type IPeople interface {
	GetName() string
}
type IPeople2 interface {
	GetName() string
	GetAge() int
}

type person struct {
	name string
}

func (p *person) GetName() string {
	return p.name
}

type person2 struct {
	name string
	age  int
}

func (p *person2) GetName() string {
	return p.name
}
func (p *person2) GetAge() int {
	return p.age
}
func main() {
	//p不可以转化为options.IPeople2接口，没有实现options.IPeople2接口中的GetAge()
	var p IPeople = &person{"jack"}
	if p2, ok := p.(IPeople2); ok {
		fmt.Println(p2.GetName(), p2.GetAge())
	} else {
		fmt.Println("p不是Ipeople2接口类型")
	}
	//p2可以转化为options.IPeople接口，因为实现了options.IPeople接口的所有方法
	var p2 IPeople2 = &person2{"mary", 23}
	if p, ok := p2.(IPeople); ok {
		fmt.Println(p.GetName())
	}

	var pp IPeople = &person{"alen"}
	if pp2, ok := pp.(*person); ok {
		fmt.Println(pp2.GetName()) //pp接口指向的对象实例是否是*person类型,*不能忘
	}
	switch pp.(type) {

	//case *person:
	//	fmt.Println("person") //判断接口的类型
	case IPeople:
		fmt.Println("options.IPeople") //判断接口的类型
		//fallthrough not applicable
	case IPeople2:
		fmt.Println("options.IPeople2")
	default:
		fmt.Println("can't found")
	}

	var ii interface{} = 43 //默认int类型
	switch ii.(type) {
	case int:
		fmt.Println("int")
	default:
		fmt.Println("can't found")
	}
}
