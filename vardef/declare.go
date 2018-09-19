package vardef

var x, y int = 1, 2
var z float32 = 3.14

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int  = 222
	b bool = true
)

var c, d int = 1, 2
var e, f = 123, "hello"
var ll = 0

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {
	g, h := 123, "world"
	//kk := 00
	println(x, y, z, a, b, c, d, e, f, g, h)
}
