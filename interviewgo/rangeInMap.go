package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	fmt.Printf("%p\n", &stus[0])
	fmt.Printf("%v\n", &stus[0])
	fmt.Printf("%p\n", &stus[1])
	fmt.Printf("%v\n", &stus[1])

	for _, stu := range stus {
		fmt.Printf(" range value &stu %p\n", &stu)
		m[stu.Name] = &stu // don't get pointer in a range loop
	}
	for k, v := range m {
		fmt.Printf("key=%s, value=%v \n", k, v)
	}

	for i := 0; i < 3; i++ {
		m[stus[i].Name] = &stus[i] // need pointer ? use index to get
	}
	for k, v := range m {
		fmt.Printf("key=%s, value=%v \n", k, v)
	}
}

func main() {
	pase_student()
}
