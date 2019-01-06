package main

import (
	"io"

	"bytes"
	"fmt"
)

const debug = true

func main() {
	var buf *bytes.Buffer
	//var buf io.Writer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		fmt.Println(buf.String())
	}

	fmt.Println(new(bytes.Buffer).WriteString("Done!\n"))

}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
