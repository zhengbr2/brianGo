package main

import (
	"fmt"
)

//interface definition
type VowelsFinder interface {
	FindVowels(str string) []rune
}

// interface wrapper
type VowelsFinderImpl func(str string) []rune

func (a VowelsFinderImpl) FindVowels(str string) []rune {
	return a(str)
}

// really implementor
func localFindVowels(ms string) []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := "Sam Andersonii, toyota"
	var v VowelsFinder = VowelsFinderImpl(localFindVowels)

	fmt.Printf("Vowels are %c", v.FindVowels(name))

}
