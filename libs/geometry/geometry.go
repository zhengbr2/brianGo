package geometry

import (
	"brianGo/libs/geometry/rectangle" // 导入自定义包
	"fmt"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))

}
