package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice :=arr[:]
	slice[0]=88
	fmt.Println( cap(slice))
	fmt.Println(arr)
	slice = append(slice, 6,7,8,9,10)
	fmt.Println( cap(slice))
	slice[1]=99
	fmt.Println(arr) // when cap(arr) expanded,means an array is created and copied
	fmt.Println(slice)


}