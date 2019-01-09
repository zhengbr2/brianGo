package main

import "fmt"

func modify1(arr *[3]int) {
	(*arr)[0] = 90
}

func modify2(arr *[3]int) {
	arr[0] = 90
}

func modify3(sls []int) {
	sls[0] = 90
}

func main() {
	//pointer to an array as a argument to a function
	m := [3]int{89, 90, 91}
	modify1(&m)
	fmt.Println("\nmodify1 using pointer to an array", m)

	n := [3]int{89, 90, 91}
	modify2(&n)
	fmt.Println("modify2 using pointer to an array", n)

	//slice as an argument to a function
	o := [3]int{89, 90, 91}
	modify3(o[:])
	fmt.Println("modify3 using slice", o)


}
