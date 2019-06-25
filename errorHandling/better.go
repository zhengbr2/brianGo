package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)


func main() {
	fmt.Println("....................erros.New....................\n")
	_, err := openFile()
	fmt.Printf("%+v", err)

	fmt.Println("\n....................erros.WithStack....................\n")
	_, err2 := openFile2()
	fmt.Printf("%+v", err2)

	fmt.Println("\n....................erros.WithMessage....................\n")
	_, err3 := openFile3()
	fmt.Printf("%+v", err3)

	fmt.Println("\n....................erros.Wrap....................\n")
	_, err4 := openFile4()
	fmt.Printf("%+v", err4)

}

func openFile() ([]byte, error) {
	return nil, errors.New("wrong file name!")
}

func openFile2() ([]byte, error) {
	_, err := os.Open("non_exist.txt")

	if err !=nil {
		err = errors.WithStack(err)    //withStack for those without stack
	}

	return []byte("bytes"), err
}

func openFile3() ([]byte, error) {
	str,err:=openFile()
	err = errors.WithMessage(err, "with self-defined error message")
	return str,err
}

func openFile4() ([]byte, error) {
	str,err:=openFile()
	err = errors.Wrap(err,"wraped msg")
	return str,err
}