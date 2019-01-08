package main

import "log"

type SyntaxError struct {
	msg    string // 错误描述
	Offset int64  // 错误发生的位置
}

func (e *SyntaxError) Error() string { return e.msg }

func Decode() *SyntaxError {
	var err *SyntaxError
	err = nil
	return err
}

func Decode2() error { // 错误，将可能导致上层调用者err!=nil的判断永远为true。
	var err *SyntaxError // 预声明错误变量
	err = nil
	return err // 错误，err永远等于非nil，导致上层调用者err!=nil的判断始终为true
}

func Decode4() error {  // 终极版
	var err error
	if (false){
		err= &SyntaxError{}
	}
	return err
}

func main() {
	s1 := Decode()
	log.Println("s1 == nil",s1 == nil)

	s2 := Decode2()
	log.Println("s2 == nil",s2 == nil)

	var s3 error =  Decode()
	log.Println("s3 == nil",s3 == nil)

	var s4 error =  Decode4()
	log.Println("s4 == nil",s4 == nil)
}
