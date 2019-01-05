package main

import "fmt"

func main() {
	ages:=[]int{6,6,6}
	ages2:=ages
	fmt.Printf("原始slice的内存地址是%p\n",ages)  //slice可以直接不取地址输出地址值， 此值为内建数组的起始位置
	fmt.Printf("原始slice的内存地址是%p\n",ages2)
	fmt.Printf("ages 的内存地址是%p\n",&ages)   // &ages , &ages2 肯定是不一样的
	fmt.Printf("ages2的内存地址是%p\n",&ages2)
	modify(ages)
	ages=append(ages, 1,2,3,4,5,6,7,8,9)
	fmt.Printf("修改后的slice的内存地址是%p\n",ages)  //内建数组的起始位置已经改变
	fmt.Printf("修改slice后， ages2的内存地址是%p\n",ages2) //不变
	fmt.Println(ages)
}

func modify(ages []int){
	fmt.Printf("函数里接收到slice的内存地址是%p\n",ages)
	ages[0]=1
}

/*
func (v Value) Pointer() uintptr {
	// TODO: deprecate
	k := v.kind()
	switch k {
	//省略无关代码
	case Slice:
		return (*SliceHeader)(v.ptr).Data
	}
}
很明显了，当是slice类型的时候，返回是slice这个结构体里，字段Data第一个元素的地址。
https://www.flysnow.org/2018/02/24/golang-function-parameters-passed-by-value.html
*/