package main

import (
	"fmt"
		"bytes"
	"os"
)

func main() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("hello"))

	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b,",","http://www.flysnow.org")
	fmt.Fprint(&b,"\n","another line from brian")
	//把Buffer里的内容打印到终端控制台
	b.WriteTo(os.Stdout)
	//b.Reset()

	b.Write([]byte("ILoveYou"))
	var p=[]byte{'a','b','c','d','e'}
	n,err:=b.Read(p[:])
	fmt.Println(n,err,string(p[:n]))   // abcde-> ILove
}
