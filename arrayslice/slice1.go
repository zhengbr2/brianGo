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
}
