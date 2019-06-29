package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func printInt(i int) {
	fmt.Printf("int i: %p\n", &i)
}
func printInt2(i *int) {
	fmt.Printf("int i: %p\n", i)
}
func printStr(s string) {
	fmt.Printf("string s: %p\n", &s)
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	data := hdr.Data
	fmt.Printf("string s data: 0x%x\n", data)
}
func printStr2(s *string) {
	fmt.Printf("string s: %p\n", s)
	hdr := (*reflect.StringHeader)(unsafe.Pointer(s))
	data := hdr.Data
	fmt.Printf("string s data: 0x%x\n", data)
}
func printSlice(s []int) {
	fmt.Printf("slice s: %p\n", &s)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	data := hdr.Data
	fmt.Printf("slice s data: 0x%x\n", data)
}
func printSlice2(s *[]int) {
	fmt.Printf("slice s: %p\n", s)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(s))
	data := hdr.Data
	fmt.Printf("slice s data: 0x%x\n", data)
}
type S struct {
	I int
}
func printStruct(s S) {
	fmt.Printf("struct s: %p, I: %p\n", &s, &(s.I))
}
func printStruct2(s *S) {
	fmt.Printf("struct s: %p, I: %p\n", &s, &(s.I))
}
func printInterface(i interface{}) {
	fmt.Printf("int i: %p\n", &i)
}
func printInterface2(i interface{}) {
	s := i.(S)
	fmt.Printf("struct s: %p, I: %p\n", &s, &(s.I))
}
func printInterface3(i interface{}) {
	s := i.(*S)
	fmt.Printf("struct s: %p, I: %p\n", s, &(s.I))
}
func main() {
	//test int
	i := 10
	fmt.Printf("int i: %p\n", &i)
	printInt(i)
	printInt2(&i)
	//test string
	s := "hello, world"
	fmt.Printf("\n\nstring s: %p\n", &s)
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	data := hdr.Data
	fmt.Printf("string s data: 0x%x\n", data)
	printStr(s)
	printStr2(&s)
	//slice and map
	sl := []int{1, 2, 3, 4, 5}
	fmt.Printf("\n\nslice s: %p\n", &sl)
	hdr2 := (*reflect.SliceHeader)(unsafe.Pointer(&sl))
	data = hdr2.Data
	fmt.Printf("slice s data: 0x%x\n", data)
	printSlice(sl)
	printSlice2(&sl)
	//struct
	ss := S{I: 10}
	ssp := &ss
	fmt.Printf("\n\nstruct s: %p, I: %p\n", ssp, &(ss.I))
	printStruct(ss)
	printStruct2(ssp)
	//interface to int
	fmt.Printf("\n\nint i: %p\n", &i)
	printInterface(i)
	//interface to struct
	fmt.Printf("\n\nstruct s: %p, I: %p\n", ssp, &(ss.I))
	printInterface2(ss)
	printInterface3(ssp)
}
