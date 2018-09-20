package pointer

import "fmt"

func main() {

	//slice := []int{0, 1, 2, 3}
	//fmt.Printf("slice: %v slice addr %p \n", slice, &slice)
	//
	//ret := changeSlice(slice)
	//fmt.Printf("slice: %v ret: %v slice addr %p \n", slice, &slice, ret)

	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Println(arr)

	slice := arr[1:4]
	slice2 := arr[2:5]

	fmt.Printf("values: arr %v, slice1 %v, slice2 %v \n", arr, slice, slice2)
	fmt.Printf("pointers: %p %p %p\n",  &arr, &slice, &slice2)

	// they refer to the same memory address

	fmt.Printf("arr[2]%p slice[1] %p slice2[0]%p\n", &arr[2], &slice[1], &slice2[0])
	fmt.Printf("arr[3]%p slice[2] %p slice2[1]%p\n", &arr[3], &slice[2], &slice2[1])

	arr[2] = 2222

	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)


	slice[1] = 1111

	fmt.Printf("arr %v, slice1 %v, slice2 %v\n", arr, slice, slice2)
}

