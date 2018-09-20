package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func SumUtil(n uint64) (result uint64) {
	if n > 0 {
		result = n + SumUtil(n-1)
		return result
	}
	return 0
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int = 10
	fmt.Printf("%d Factorial is: %d\n", i, Factorial(uint64(i)))
	fmt.Printf("%d SumUtil is: %d\n", i, SumUtil(uint64(i)))
	fmt.Printf("%d fibonacci is: %d\n", i, fibonacci(i))
}
