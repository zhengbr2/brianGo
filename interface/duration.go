package main

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

type Duration int64

func (a Duration) String() string { //toString() in golang

	var b bytes.Buffer
	fmt.Fprint(&b, int64(a), "(my local duration)")
	return (b.String())
}

func main() {
	{
		var i time.Duration = 111
		var j int64 = 222
		fmt.Println(i, j)
		j = int64(i)
		fmt.Println(i, j) // same as below
		fmt.Println(i.String(), j)
	}

	{
		var i Duration = 111
		var j int64 = 222
		fmt.Println(i, j)    // String()值接收，如果是引用接受， 这里就不会调用自定义输出
		j = int64(i)
		fmt.Println(&i, j) // can use ref also
		fmt.Println(i.String(), j)
		log.Println()
	}

}
