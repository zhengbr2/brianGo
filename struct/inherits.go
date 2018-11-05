package main

import "fmt"

type Hello interface {
	hello()
}

type user struct {
	name  string
	email string
}
func (u user) sayHello() {
	fmt.Println("Hello，i am a user")
}
func (u user) hello() {
	fmt.Println("Hello，i am a user")
}

type human struct {
	name  string
	email string
}

func (u human) sayHello() {
	fmt.Println("Hello，i am a human")
}
func (u human) hello() {
	fmt.Println("Hello，i am a human")
}


type admin struct {
	user
	level string
	human
}
func (a admin) sayHello() {
	fmt.Println("Hello，i am a admin")
}


func sayHello(h Hello) {
	h.hello()
}

func main() {
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员", human{"brian", "zb@dd.org"}}
	ad.user.sayHello()
	ad.sayHello()
	ad.human.sayHello()

	sayHello(ad.user) //使用user作为参数
	//sayHello(ad)//使用admin作为参数 //ambiguous
	sayHello(ad.human) //使用human作为参数
}



