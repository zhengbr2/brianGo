package main

import "fmt"

func main(){
	num := 10
	switch  {  //省略表达式
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
		fallthrough
	case num >= 0 && num <= 100:
		fmt.Println("num is greater than 0 and less than 100")

	case num >= 0:
		fmt.Println("num is greater than 0")
	case num >= 100:
		fmt.Println("num is greater than 100")
	default:

	}

}

