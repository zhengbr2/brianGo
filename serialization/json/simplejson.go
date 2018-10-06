package main

import (
	"fmt"
	sj "github.com/bitly/go-simplejson"
)

func main() {
	js, err := sj.NewJson([]byte(`{
	"test": {
		"array": [1, "2", 3],
		"int": 10,
		"float": 5.150,
		"bignum": 9223372036854775807,
		"string": "simplejson",
		"bool": true
	}
}`))
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	bignum, _ := js.Get("test").Get("bignum").Int64()
	fmt.Println(err, arr, i, ms, bignum)
}
