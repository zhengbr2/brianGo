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