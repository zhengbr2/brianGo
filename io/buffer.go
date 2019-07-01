package main

import (
	"io"

	"bytes"
	"fmt"
)

const debug = false

func main() {
	var buf *bytes.Buffer
	//var buf io.Writer  // 这里应该定义接口
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
