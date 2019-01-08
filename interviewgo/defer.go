package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("index 1:", a, b)  // defer 获得参数后， 不会受后来修改的影响
	a = 6
	b = 7
	defer calc("index 2:", a, b)

}
