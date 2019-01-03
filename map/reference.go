package main

import (
	"fmt"
	)
func main() {
	persons:=map[string]int{"brian":19,"quan":20}
	p2:=persons
	p3:=persons
	fmt.Printf("原始map的内存地址是：%p\n",&persons)
	fmt.Printf("原始map的内存地址是：%p\n",&p2)
	fmt.Printf("原始map的内存地址是：%p\n",&p3)
	modify(persons)
	fmt.Println("map值被修改了，新值为:",persons)
}

func modify(p map[string]int){
	fmt.Printf("函数里接收到map的内存地址是：%p\n",&p)
	p["张三"]=20
}
