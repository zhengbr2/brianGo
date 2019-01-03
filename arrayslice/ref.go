package main

import "fmt"

func main() {
	ages:=[]int{6,6,6}
	ages2:=ages
	fmt.Printf("原始slice的内存地址是%p\n",ages)
	fmt.Printf("原始slice的内存地址是%p\n",ages2)
	modify(ages)
	fmt.Println(ages)
}

func modify(ages []int){
	fmt.Printf("函数里接收到slice的内存地址是%p\n",ages)
	ages[0]=1
}
