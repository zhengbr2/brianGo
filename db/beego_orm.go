package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"time"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	//注册定义的model

	orm.RegisterModel(new(User))
	// 创建table
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id         int `PK` //如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
	Username   string
	Departname string
	Created    time.Time
}

func main() {
	o := orm.NewOrm()
	var user User
	user.Username = "brian"
	user.Departname = "Dev"
	user.Created = time.Now()

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}

	// 更新表
	user.Username = "ivyfan"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
