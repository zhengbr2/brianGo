package main

import (
	"os"
	"io"
	"bytes"
	"fmt"
		"syscall"
	"errors"
	)

func main(){
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // success: f == os.Stdout
	c ,okay:= w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	fmt.Println(f,c,okay)

	{
		_, err := os.Open("/no/such/file")
		fmt.Println(err) // "open /no/such/file: No such file or directory"
		fmt.Printf("%#v\n", err)
	}
	{
		_, err := os.Open("/no/such/file")
		fmt.Println(IsNotExist(err)) // "true"
	}
}


type PathError struct {
	Op string
	Path string
	Err error
}
func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

var ErrNotExist = errors.New("file does not exist")
// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}