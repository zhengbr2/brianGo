package main

import (
	"database/sql"
	"fmt"
	//"time"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("brian", "development", "2006-04-09")
	checkErr(err)

	{
		stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
		checkErr(err)

		res, err := stmt.Exec("dindin", "kindergarten", "2012-03-07")
		checkErr(err)
		id, err := res.LastInsertId()
		log.Println("last insert id: ", id)
	}

	id, err := res.LastInsertId()
	checkErr(err)

	log.Println("last insert id:", id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("mingkun", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	log.Println("affected number:", affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	log.Println("affected number after delete:", affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
