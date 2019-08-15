
package main

import "fmt"

type Hello interface {
	hello()
}

type user struct {
	name  string
	email string
}


func (u user) hello() {
	fmt.Println("Hello，i am a user")
	u.bye()
}

func (u user) bye() {
	fmt.Println("bye() user")
}



type admin struct {
	user
	level string
	//human
}


func (u admin) bye() {
	fmt.Println("bye() admin")
}



func main() {
	ad := admin{user{"张三", "zhangsan@flysnow.org"}, "管理员"}

	ad.hello() // doesn't dispatch to admin.bye()

}
