package main

import (
	"fmt"
	"unsafe"
)
func main() {
	persons:=map[string]int{"brian":19,"quan":20}
	p2:=persons
	p3:=persons
	fmt.Printf("原始map的内存地址是：%p\n",&persons)
	fmt.Printf("p2的内存地址是：%p\n",&p2)
	fmt.Printf("p3的内存地址是：%p\n",&p3)
	fmt.Printf("persons：%d\n",unsafe.Sizeof(persons))
	fmt.Printf("&persons：%d\n",unsafe.Sizeof(&persons))
	fmt.Printf("p2：%d\n",unsafe.Sizeof(p2))
	fmt.Printf("p3：%d\n",unsafe.Sizeof(p3))

	fmt.Printf("注意这个：persons的数组地址: %p:\n",persons)
	modify(persons)
	modify(persons)

	fmt.Println("map值被修改了，新值为:",persons)
}

func modify(p map[string]int){
	fmt.Printf("函数里接收到map的内存地址是：%p\n",&p)
	fmt.Printf("注意这个：函数里接收到map的内存地址是：%p\n",p)
	p["张三"]=20
}
