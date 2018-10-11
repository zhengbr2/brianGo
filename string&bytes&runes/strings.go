package main

import "fmt"
import (
	s "strings"
	"unicode/utf8"
)

var p = fmt.Println

func main() {

	 var str string
	 if str == ""{
	 	str = "default"
	 	println(str)
	 }

	x := "text"
	fmt.Println(x[0]) //print 116
	fmt.Printf("%T\n",x[0]) //prints uint8

	data := "♥"
	fmt.Println(len(data))  //3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	data = "é"
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
}
