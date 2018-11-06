package main

import "log"

type SyntaxError struct {
	msg    string // 错误描述
	Offset int64  // 错误发生的位置
}

func (e *SyntaxError) Error() string { return e.msg }

func Decode() *SyntaxError { // 错误，将可能导致上层调用者err!=nil的判断永远为true。
	var err *SyntaxError // 预声明错误变量
	if false {
		err = &SyntaxError{}
	}
	return err // 错误，err永远等于非nil，导致上层调用者err!=nil的判断始终为true
}

func main() {
	s := Decode()
	log.Println(s == nil)
}
