package main

import (
	"fmt"
	"github.com/pkg/errors"
)


func main() {
	conent, err := openFile2()
	if err != nil {
		fmt.Printf("%+v", err)
	} else {
		fmt.Println(string(conent))
	}
}

func openFile() ([]byte, error) {
	return nil, errors.New("wrong file name!")
}

func openFile2() ([]byte, error) {
	conent, err := openFile()
	//if err !=nil {
	//	err = errors.WithMessage(err, "failed in openFile2")   // one more message
	//}

	//if err !=nil {
	//	err = errors.WithStack(err)    //withStack for those without stack
	//}

	if err !=nil {
		err = errors.Wrap(err, "msg in wrap()")   // one more message + stack
	}
	return conent, err
}