package main

import "fmt"

func main() {
	isSpace := func(ch byte) bool {
		switch ch {
		case ' ': //error

		case '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t')) //prints true (ok)
	fmt.Println(isSpace(' '))  //prints false (not ok)

	data := []int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])
}
