package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer

	//使用Write方法为写入字符串
	bn, be := b.Write([]byte("hello"))
	log.Println("number of bytes wrritten:", bn)
	if be != nil {
		log.Fatal(be)
	}
	fmt.Println(b.String())

	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b, ",", "http://www.flysnow.org\n")

	//把Buffer里的内容打印到终端控制台
	fmt.Println("1:",b.String())
	b.WriteTo(os.Stdout)   // 清空buffer 从头写
	fmt.Println("2:",b.String())
	bn, be = b.Write([]byte("ILoveYou"))
	fmt.Println("3:",b.String())
	log.Println("number of bytes wrritten:", bn)

	log.Println("current length of b:", b.Len())
	if be != nil {
		log.Fatal(be)
	}
	fmt.Println(b.String())

	var p = []byte{'a', 'b', 'c', 'd', 'e'}
	n, err := b.Read(p[:])
	log.Println(n, err, string(p[:n])) // abcde-> ILove
	n, err = b.Read(p[:])
	log.Println(n, err, string(p[:n])) // abcde-> You
	n, err = b.Read(p[:])
	log.Println(n, err, string(p[:n])) // _ 0 EOF

	data, err := ioutil.ReadAll(&b)
	fmt.Println(string(data), err)

}
