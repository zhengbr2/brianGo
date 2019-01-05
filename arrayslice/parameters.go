package main

import "fmt"

func find(num int, nums ...int) {

	//以下证明nums是slice
	fmt.Printf("type of nums is %T\n", nums)
	//nums = append(nums, 1)
	_= len(nums)

	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	nums := []int{89, 90, 95}
	find(89, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	welcome[0] = "nice"
	fmt.Println(welcome)

}
