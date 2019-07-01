package main

import (
	"bytes"
	"fmt"
	"os"
)

func main(){
	write()
	fmt.Println(".........................................")
	read()
}

func write(){
	s := []byte(" world")
	buf := bytes.NewBufferString("hello")
	fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
	buf.Write(s) //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
	buf.WriteString(", what the fuck") //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
	buf.WriteByte('!') //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
	buf.WriteRune('好')
	fmt.Printf("%v,%v\n",buf.String(),buf.Bytes())
	file,_ := os.Create("test.txt")
	buf.WriteTo(file)
	file.Close()
}

func read(){
	s1 := []byte("hello")
	buff :=bytes.NewBuffer(s1)
	fmt.Println(buff.String())  //hello
	s2 := []byte(" world")
	buff.Write(s2)
	fmt.Println(buff.String())  //hello world

	s3 := make([]byte,3)
	buff.Read(s3)
	fmt.Println(string(s3))  //hel,s3的容量为3，只能读3个
	fmt.Println(buff.String()) //lo world

	buff.Read(s3)  // 会把s3覆盖掉
	fmt.Println(string(s3))  // lo
	fmt.Println(buff.String())  // world
	b,err:=buff.ReadByte()
	fmt.Println(string(rune(b)),err)  // w <nil>
	fmt.Println(buff.String())  // orld

	file, _ := os.Open("test.txt")
	buf := bytes.NewBufferString("!!!")
//	buf=bytes.NewBuffer([]byte{})
	buf.ReadFrom(file)
	fmt.Println(buf.String()) //bob hello world
	file.Close()
}