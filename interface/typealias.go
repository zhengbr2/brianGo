package main

import (
	"fmt"
	"unsafe"
)

//interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

//MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	var v VowelsFinder
	fmt.Printf("length of interface VowelsFinder: %d\n", unsafe.Sizeof(v))
	fmt.Println(v==nil)
	v = name // possible since MyString implements VowelsFinder
	fmt.Println(unsafe.Sizeof(name)) //16, value + type
	fmt.Println(unsafe.Sizeof( new (int32)))  //8
	fmt.Printf("Vowels are %c", v.FindVowels())

}
