package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var count = 0

func greeting(w http.ResponseWriter, r *http.Request) {

	count = count + 1
	fmt.Fprintf(w, "Hello World:%d __", count) //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", greeting) //设置访问的路由
	http.HandleFunc("/cookie", plantCookie)
	http.HandleFunc("/readcookie", readCookie)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func plantCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "zhengbr2", Expires: expiration}
	http.SetCookie(w, &cookie)
	w.Write([]byte("hellow world!"))
}
func readCookie(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	for index, ck := range cookies {
		log.Println("cookie index", index, "content:", ck)

	}

}
