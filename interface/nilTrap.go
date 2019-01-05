package main

import (
	"github.com/smallnest/rpcx/log"
	"strconv"
)

type myError struct {
	errnum int
}

func (e *myError) Error() string {
	return "error number is :" + strconv.Itoa(e.errnum)
}

func main() {

	var s string
	println("s ==\"\":", s == "")

	var e myError
	checkErr(&e)

	// use below
	//var e error
	//checkErr(e)

}

func checkErr(e error) {
	if e != nil {
		log.Info("error deteced!")
	} else {
		log.Debug("no error")
	}

}
