package main

import (
	ct "context"
	"github.com/gorilla/context"
	"net/http"
	"strconv"
)

func main() {
	//启动一个Web服务
	http.Handle("/", http.HandlerFunc(myHander))
	http.Handle("/v2", http.HandlerFunc(myHander2))
	http.Handle("/v3", context.ClearHandler(http.HandlerFunc(myHander)))
	http.Handle("/v4", http.HandlerFunc(myHander4)) // replaced by go 1.7 context

	http.ListenAndServe(":1234", nil)
}

//定义一个Hander
func myHander(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	context.Set(r, "user", "张三")
	context.Set(r, "age", 18)

	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHander(rw, r)
}

func doHander(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	user := context.Get(r, "user").(string)
	age := context.Get(r, "age").(int)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}

func myHander2(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	context.Set(r, "user", "张三")
	context.Set(r, "age", 18)

	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHander2(rw, r)
}

func doHander2(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	allParams := context.GetAll(r)
	user := allParams["user"].(string)
	age := allParams["age"].(int)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))
}

func myHander4(rw http.ResponseWriter, r *http.Request) {
	//模拟为Request附加值，这里附加了2个
	userContext := ct.WithValue(ct.Background(), "user", "张三")
	ageContext := ct.WithValue(userContext, "age", 18)
	rContext := r.WithContext(ageContext)

	//这个模拟一个方法或者函数的调用，大部分情况下可能不在一个包里
	doHander4(rw, rContext)
}

func doHander4(rw http.ResponseWriter, r *http.Request) {
	//我们从这个Request里取出对应的值。
	user := r.Context().Value("user").(string)
	age := r.Context().Value("age").(int)

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("the user is " + user + ",age is " + strconv.Itoa(age)))

}
