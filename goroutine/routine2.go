package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c_%d ", i,i)
		//fmt.Printf("%d ", i)
	}
}

func chars() {
	for i := 96; i <= 105; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("rune %s \n ", string(rune(i)))

	}
}

func main() {
	go numbers()
	go alphabets()
	//go chars()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\nmain terminated")
}
