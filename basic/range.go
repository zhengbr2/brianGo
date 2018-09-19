package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)


	sum = 0
	for  num := range nums {
		sum += num   // sum index :)
	}
	fmt.Println("sum:", sum)


	sum = 0
	for i, num := range nums {
		sum += num
		if num == 3 {
			fmt.Println("index:", i)
		}
		if i == 1 {
			fmt.Println("sum:", sum)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len %d, cap %d,  s:%v\n", len(s), cap(s), s)
}
