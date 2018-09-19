package class

import (
	"fmt"
	_ "brianGo/libs/mymath"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 4
	fmt.Println("Area of Circle(c1) = ", c1.getArea())
	fmt.Println("Perimeter of Circle(c1) = ", c1.getPerimeter())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func (c Circle) getPerimeter() float64 {

	return 3.14 * c.radius * 2.0
}
