package main

import "fmt"

func main(){

	arr1 := [...] int {0,1,2,3,4}
	sli0 :=arr1[:]
	sli1 := arr1[1:]
	sli2 := arr1[2:]
	fmt.Printf("the adderss arr, arr1, arr2: %p, %p , %p\n", &arr1,&arr1[1], &arr1[2])
	fmt.Printf("the adderss of slice pointer: %p, %p, %p \n", sli0,sli1,sli2)
	fmt.Printf("the adderss of slice self: %p, %p, %p \n", &sli0,&sli1,&sli2)

	sli0 = append(sli0, 5,6,7,8,9,10)
	fmt.Printf("the adderss after reslice: %p\n",sli0)

	{
		arr:=[]int{1,2,3,4,5}
		slice:=arr[1:2]
		fmt.Println(slice)
		slice =append(slice,6,7,8) //不超过 cap 就覆盖 arr, 超过 cap 新建底层数组反而没事
		fmt.Println(slice)
		fmt.Println(arr)
	}
}
