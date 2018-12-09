package main

import "fmt"

func main(){

	arr := [8] int {1,2,3,4,5,6,7,8}
	sli := arr[0:5]
	fmt.Println(sli)
	sli=append(sli, 9)   //base arr is updated!
	fmt.Println(sli)
	fmt.Println(arr)
	sli=append(sli, 10,10,10,10)  // once over capacity, make a new slice and the old one not updated
	// we shall not append into a slice
	fmt.Println(sli)
	fmt.Println(arr)

	s := make([]int, 2)
	mdSlice(s)
	fmt.Println("append into 0 capacity:",s)

	s = make([]int, 4,4)
	mdSlice(s)
	fmt.Println("append into 4 capacity:",s)

	arr2:=[4]int{}
	s1 := arr2[:1]
	s2 := append(s1, 2)

	fmt.Println("append into :",s1,s2)
	//可以通过slice 改变array，但是不一定能改变其他的slice，因为其Len已经指定了。
}

func mdSlice(s []int) {
	s = append(s, 1)
	s = append(s, 2)
}