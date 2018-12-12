package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	//w = time.Second         // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	//rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	fmt.Println(w, rwc)

	os.Stdout.Write([]byte("hello")) // OK: *os.File has Write method
	os.Stdout.Close()                // OK: *os.File has Close method

	w = os.Stdout
	w.Write([]byte("hello")) // OK: io.Writer has Write method
	//w.Close()                // compile error: io.Writer lacks Close method

	var w2 io.Writer = new(bytes.Buffer)
	var w3 io.Writer = (*bytes.Buffer)(nil)
	fmt.Println(w2, w3)


}

func main2() {
	var b bytes.Buffer
	fmt.Fprint(&b,"Hello World")
	fmt.Println(b.String())
}
