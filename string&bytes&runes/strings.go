package main

import "fmt"
import (
	s "strings"
	"unicode/utf8"
)

var p = fmt.Println

func main() {

	var str string
	//if str==nil  str by defalt is "" and not comparable to nil

	if str == "" {
		str = "default"
		println(str)
	}

	x := "text"
	slice := []byte(x)
	slice[0] = 'T'
	println(x, ":", string(slice)) //text:Text
	fmt.Println(x[0])              //print 116
	fmt.Println(slice[0])              //print 86
	fmt.Printf("%T\n", x[0])       //prints uint8

	//data := "♥"
	//fmt.Println(len(data))                    //3
	//fmt.Println(utf8.RuneCountInString(data)) //prints: 1


	data := "é"
	// not data ="e"  !!
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
	p("value is:", ("hello" == "hello"))

	var c = "hello"[1]
	fmt.Printf("type of hello[1] %T",c)   //unit8
}
