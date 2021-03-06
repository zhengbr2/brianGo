package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	//write 不断地添加到缓冲区, 底层slice不断增长，
	write()
	fmt.Println(".........................................")
	//读完了就没有了，drain了
	read()
}

func testEmpty() {
	bf := bytes.Buffer{}
	bf.Write([]byte("quan"))
	fmt.Println(bf)
	bt := make([]byte, 4)
	bf.Read(bt)
	fmt.Println(bf)
	bf.Read(bt)
}

func write() {
	s := []byte(" world")
	buf := bytes.NewBufferString("hello")
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())
	buf.Write([]byte("!!")) //buf will grow if capacity no enough
	buf.Write(s)            //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())
	buf.WriteString(", what the fuck") //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())
	buf.WriteByte('!') //buf will grow if capacity no enough
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())
	buf.WriteRune('好')
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes())
	file, _ := os.Create("test.txt")
	buf.WriteTo(file)
	fmt.Printf("%v,%v\n", buf.String(), buf.Bytes()) // the buffer is drained.
	file.Close()
}

func read() {
	s1 := []byte("hello")
	buff := bytes.NewBuffer(s1)
	fmt.Println(buff.String()) //hello
	s2 := []byte(" world")
	buff.Write(s2)
	fmt.Println(buff.String()) //hello world

	s3 := make([]byte, 3)
	buff.Read(s3)
	fmt.Println(string(s3))    //hel,s3的容量为3，只能读3个
	fmt.Println(buff.String()) //lo world
	fmt.Println(buff.String()) //lo world

	buff.Read(s3)              // 会把s3覆盖掉
	fmt.Println(string(s3))    // lo
	fmt.Println(buff.String()) // world
	b, err := buff.ReadByte()
	fmt.Println(string(rune(b)), err) // w <nil>
	fmt.Println(buff.String())        // orld

	file, _ := os.Open("test.txt")
	buf := bytes.NewBufferString("!!!")
	//	buf=bytes.NewBuffer([]byte{})
	buf.ReadFrom(file)
	fmt.Println(buf.String()) //bob hello world
	file.Close()

	buff = bytes.NewBufferString("Hello!!!")
	s4 := []byte("hello")
	n, err := buff.Read(s4)
	fmt.Println(string(s4[0:n])) // Hello
	fmt.Println(buff.String())   // !!!
	n, err = buff.Read(s4)
	fmt.Println(buff.String())   // ""
	fmt.Println(string(s4[0:n])) // !!!

	fmt.Println("empty now")
	buff.Read(s4)              //buff drained
	fmt.Println(buff.String()) // ""
	buff.Reset()               //buff drained
	fmt.Println(buff.String()) // ""

}
