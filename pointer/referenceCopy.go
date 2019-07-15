package main

import "fmt"

func refSwap(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func refSwap2(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前，a, b 的值 : %d,  %d\n", a, b)

	refSwap(&a, &b)

	fmt.Printf("交换1，a, b 的值 : %d,  %d\n", a, b)

	refSwap2(&a, &b)

	fmt.Printf("交换2，a, b 的值 : %d,  %d\n", a, b)

}
