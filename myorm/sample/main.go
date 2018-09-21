package main

import "fmt"
import (
	"brianGo/myorm"
	"time"
)

type User struct {
	ID        int64     `json:"id"`         // 自增主键
	Age       int64     `json:"age"`        // 年龄
	FirstName string    `json:"first_name"` // 姓
	LastName  string    `json:"last_name"`  // 名
	Email     string    `json:"email"`      // 邮箱地址
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

func main() {
	ormDB, _ := myorm.Connect("root:root@tcp(127.0.0.1:3306)/ormdb")
	users := myorm.Table(ormDB, "user")

	a := users()
	fmt.Println(a)

	user1 := User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Cat",
	}
	user2 := User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Curise",
	}
	user3 := User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Hanks",
	}
	user4 := map[string]interface{}{
		"age":        30,
		"first_name": "Tom",
		"last_name":  "Zzy",
	}
	users().Insert([]interface{}{user1, user2})
	users().Insert(user3)
	users().Insert(user4)
}
