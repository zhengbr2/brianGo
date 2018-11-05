package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
	r.length = 88
}

func (r *rectangle) area2() {
	fmt.Printf("Area Method *result: %d\n", (r.length * r.width))
	r.length = 88
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func area(r rectangle) {

	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
	r.length = 99
}

func main() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("1) parameters: %d, %d\n", r.length, r.width)
	area(r)
	fmt.Printf("2) parameters: %d, %d\n", r.length, r.width)
	r.area()
	fmt.Printf("3) parameters: %d, %d\n", r.length, r.width)

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p.area() //通过指针调用值接收器
	fmt.Printf("4) parameters: %d, %d\n", r.length, r.width)

	p.area2()                                                //通过指针调用值接收器
	fmt.Printf("5) parameters: %d, %d\n", r.length, r.width) //changed!!!!

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
