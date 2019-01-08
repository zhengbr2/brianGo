package main

import "fmt"

type T struct {
	Name string
}

func (t T) M1() {
	t.Name = "name1"
}

func (t *T) M2() {
	t.Name = "name2"
}


type Intf interface {
	M1()
	M2()
}

type S struct {
	T
}

type SP struct {
	*T
}

func main(){
	{
		fmt.Println(" ---------- senario 1 , no interface------------------")
		// 如果不涉及到 interface , 那么用值调用或者指针调用没有任何区别， 能否修改，取决于方法声明时是否指定为指针接受者。
		t1 := T{"t1"}
		fmt.Println("M1调用前：", t1.Name)
		t1.M1()
		fmt.Println("M1调用后：", t1.Name)
		fmt.Println("M2调用前：", t1.Name)
		t1.M2()   //changed
		fmt.Println("M2调用后：", t1.Name)


		t2 := &T{"t2"}
		fmt.Println("M1调用前：", t2.Name)
		t2.M1()
		fmt.Println("M1调用后：", t2.Name)
		fmt.Println("M2调用前：", t2.Name)
		t2.M2()
		fmt.Println("M2调用后：", t2.Name)
	}

	{
		fmt.Println(" ---------- senario 2, assgine to interface ------------------")
		var t1 T = T{"t1"}
		var t2 Intf
		// 如果赋值给接口，就不一样了
		//t2 = t1   //T does not implement Intf (M2 method has pointer receiver)
		t2 = &t1
		t2.M1()
		t2.M2()
		fmt.Println("Intf调用后：", t1.Name)
	}

	{
		fmt.Println(" ---------- senario 3 , embed struct by value ------------------")
		t1 := T{"t1"}
		s := S{t1}   // value copy here
		fmt.Println("M1调用前：", s.Name)
		s.M1()
		fmt.Println("M1调用后：", s.Name)
		fmt.Println("M2调用前：", s.Name)
		s.M2()
		fmt.Println("M2调用后：", s.Name)
		fmt.Println(t1.Name)   // t not changed!

		//var intf Intf = s 编译错误
		var intf Intf = &s
		intf.M1()
		intf.M2()   // 改变s.T.Name, t1 不变
		fmt.Println("intf.M2()调用后：", s.Name,t1.Name)

	}
	{
		fmt.Println(" ---------- senario 4 , embed struct by pointer ------------------")
		t1 := T{"t1"}
		s := SP{&t1}
		fmt.Println("M1调用前：", s.Name)
		s.M1()
		fmt.Println("M1调用后：", s.Name)
		fmt.Println("M2调用前：", s.Name)
		s.M2()
		fmt.Println("M2调用后：", s.Name)
		fmt.Println(t1.Name)   //s.T.Name 其实 和 .Name一样的地址！ 都变了


		t1 = T{"t1"}
		s = SP{&t1}
		var intf Intf = s
		//var intf2 Intf = &s   两者运行都一样的输出结果， 见 senario1
		intf.M1()
		intf.M2()
		fmt.Println("intf.M2()调用后：", s.Name)
		fmt.Println(s.Name, t1.Name)

	}

}


