package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 // before we can run this , arr is initialized as 0.0.0.0..
	}
	sl := append(n[:], 888)
	fmt.Printf("after append 888 %v\n", sl)

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	// not able to print n[11]

	{
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a // a copy of a is assigned to b
		b[0] = "Singapore"
		fmt.Println("size of b is ", unsafe.Sizeof(b))
		fmt.Println("size of abcdefgasdfasdf is ", unsafe.Sizeof("abcdefgasdfasdf"))   // size if 16, 8bytes for address, 8bytes (int64) for str length
		fmt.Println("a is ", a)
		fmt.Println("b is ", b)
		fmt.Println(len(b), cap(b))
		var c [5]string
		fmt.Println(len(c), cap(c))

		fmt.Printf("c[1]=%s", c[1]) // array can be initiliazed as ""
	}

	fmt.Printf("---------------------------------\n")
	{
		a := []string{"USA", "China", "India", "Germany", "France"} // this is a slice
		b := a                                                      // a copy of a is assigned to b
		b[0] = "Singapore"
		fmt.Println("a is ", a)
		fmt.Println("b is ", b) //both changed.
		fmt.Println("size of slice ", unsafe.Sizeof(b))

	}
	{
		array := [5]*int{1: new(int), 3:new(int)}
		*array[1] = 1
		*array[3] = 3
		//*array[2] = 2
		println(array[1])
		println(array[2])   // point the memory address
		println(*array[1])
		// println(*array[2]) error
	}

}
