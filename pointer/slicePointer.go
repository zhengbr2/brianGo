package main

import "fmt"

func main() {


	arr := [5]int32{0, 1, 2, 3, 4}
	fmt.Println(arr)

	slice := arr[1:4]
	slice2 := arr[2:5]

	fmt.Printf("values: arr %v, slice1 %v, slice2 %v \n", arr, slice, slice2)
	fmt.Printf("pointers address 1: %p %p %p\n", &arr, &slice, &slice2)
	fmt.Printf("pointers address 2:  %p %p %p\n", &arr, slice, slice2)   // slice 起始第一位

	// they refer to the same memory address

	fmt.Printf("arr[2]%p slice[1] %p slice2[0]%p\n", &arr[2], &slice[1], &slice2[0])
	fmt.Printf("arr[3]%p slice[2] %p slice2[1]%p\n", &arr[3], &slice[2], &slice2[1])

// 没有用append 不破坏3个arr/slice 同用一个底层数组的内在
	arr[2] = 2222
	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)

	slice[1] = 1111
	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)
}
