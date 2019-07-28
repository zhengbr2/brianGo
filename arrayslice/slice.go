package main

import (
	"fmt"
	"unsafe"
	"math/rand"
)

func main() {

	ranSlice:=make([]int, 3, rand.Intn(100))
	fmt.Println("randon slice len, and cap", len(ranSlice),cap(ranSlice))


	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	fmt.Printf("type of array %T, size of array related to lengh! %d\n", arr, unsafe.Sizeof(arr))

	sl := arr[0:3]
	//slb := arr[0:3]
	//fmt.Print(sl == slb)  not comparable
	//sl2 := sl
	//sl2[0] == 99
	fmt.Println(sl)
	fmt.Printf("type of slice %T, size of slice related to lengh! %d\n", sl, unsafe.Sizeof(sl))
	fmt.Printf("type of slice %T, size of slice related to lengh! %d\n", sl, unsafe.Sizeof(arr[0:7]))  // 地址， 长度， 容量 3×8=24
	fmt.Println(arr)

	var s []int
	s = append([]int(nil), 1, 2, 3)
	//s = append(nil, 1,2,3)  // compilation error

	s = make([]int, 5, 10)
	printSlice(s)
	s[0] = 111
	s[1] = 222
	s[2] = 333
	s[3] = 444
	s[4] = 555
	//	s[5] = 666
	printSlice(s)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{4: 1}
	printSlice(slice1)
	printSlice(slice2)

	array := [5]int{4: 1}
	printSlice(array[:])
	array2 := [...]int{4: 1}
	printSlice(array2[:])

	var ok = append(s, 666)
	//var ok = append(s, []int{1,,1}) not accepted
	var ok2 = append(s, 777, 888, 888, 0, 999, 99)

	printSlice(ok)
	printSlice(ok2)
	var s2 [10]int
	printSlice(s2[:])
	var s3 []int
	printSlice(s3)

	var s4 = []int{}
	printSlice(s4)
	//numbers := []int{0,1,2,3,4,5,6,7,8}
	//printSlice(numbers)
	//fmt.Println("numbers ==", numbers)
	//fmt.Println("numbers[1:4] ==", numbers[1:4])
	///* 默认下限为 0*/
	//fmt.Println("numbers[:3] ==", numbers[:3])
	//
	///* 默认上限为 len(s)*/
	//fmt.Println("numbers[4:] ==", numbers[4:])
	//
	//numbers1 := make([]int,0,5)
	//printSlice(numbers1)

	fmt.Println("-----------------------------")
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 2, 3)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	printSlice(numbers1)

	fmt.Println("-----------------------------")
	var numbers3 []int
	numbers3 = append(numbers3, 1, 2, 3, 4, 5)
	printSlice(numbers3)
	numbers3 = append(numbers3, 1, 2, 3)
	printSlice(numbers3)

	fmt.Println("-----------------------------")
	{
		var from = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		var to = []int{9, 9, 9, 9, 9}

		cnt:=copy(to, from)
		printSlice(from)
		printSlice(to)
		fmt.Println(cnt)

	}
	{
		var from = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		var to = []int{9, 9, 9, 9, 9}

		copy(from, to)
		printSlice(from)
		printSlice(to)
	}
	{
		var temp = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
		s1 := temp[0:5]
		s2 := temp[3:8]
		printSlice(s1)
		printSlice(s2)
		copy(s1, s2)
		printSlice(temp)
		printSlice(s1)
		printSlice(s2)
		s1=s1[0:len(s1)+3]   //len 8, cap:9
		printSlice(s1)

	}

	{
		s0:=make([]int, 0)
		s1:=make([]int, 5)
		s2:=make([]int, 0, 5)
		fmt.Println(s0)
		fmt.Println(s1)
		fmt.Println(s2)
	}

}

func printSlice(s []int) {
	fmt.Printf("len %d, cap %d,  s:%v\n", len(s), cap(s), s)
}
