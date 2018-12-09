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
	fmt.Println(sli)
	fmt.Println(arr)
}
