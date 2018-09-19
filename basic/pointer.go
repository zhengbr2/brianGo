package main

import "fmt"

func main() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("a 的ptr为  %d\n", ptr)
	fmt.Printf("a 的ptr为  %x\n", ptr)
	fmt.Printf("*ptr 为 %d\n", *ptr)

	fmt.Printf("变量的地址: %x\n", &a)

	var nptr *int

	fmt.Printf("nptr 的值为 : %x\n", nptr)
	if nptr == nil {
		fmt.Printf("null pointer")
	}

}
