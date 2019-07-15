package main

import "fmt"

func main() {
	var m map[int64]int64
	m = make(map[int64]int64, 1)

	// map 增加元素后指向的首地址依然不变
	fmt.Printf("m 原始的地址是：%p\n", m)  // 0xc000062240
	changeM(m)
	fmt.Printf("m 改变后地址是：%p\n", m)  // 0xc000062240
	fmt.Println("m 长度是", len(m))
	fmt.Println("m 参数是", m)

}
	// 改变map的函数
	func changeM(m map[int64]int64) {
		fmt.Printf("m 函数开始时地址是：%p\n", m) // 0xc000062240
		var max = 5
		for i := 0; i < max; i++ {
			m[int64(i)] = 2
		}
		fmt.Printf("m 在函数返回前地址是：%p\n", m) // 0xc000062240
	}


