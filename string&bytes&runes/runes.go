package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Println("Printing Bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printCharsUsingFor(p string) {
	fmt.Println("Printing runes using for")
	s:=p
	runes := []rune(s)
	println(&s, &runes) // diff address
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytesUsingForRange(s string) {
	fmt.Println("Printing characters and bytes using for Range")
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}


func mutate2(s []rune) []rune {
	fmt.Println("String mutate")
	s[0] = 'a'
	return s
}

func main() {

	{

	}
	name := "Señor"
	printBytes(name)
	fmt.Println("\n")

	printCharsUsingFor(name)
	fmt.Println("\n")

	printCharsAndBytesUsingForRange(name)
	fmt.Println("\n")

	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println("String constructed from hexadecimal byte slice")
	fmt.Println(str)
	fmt.Println("length", len(str))         // len is 5
	fmt.Println("length", len([]rune(str))) // len is 4

	bs:=[]byte("é")
	fmt.Println("bs[0]",bs[0])  //195
	fmt.Println("bs[1]",bs[1])  //169
	r3 := rune('é')
	fmt.Println(string(r3))

	byteSlice = []byte{67, 97, 102, 195, 169} //decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println("String constructed from decimal byte slice")
	fmt.Println(str + "\n")

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println("String constructed from rune slice")
	fmt.Println(str)
	fmt.Println()

	length(str)
	fmt.Println(len(str)) //6

}
