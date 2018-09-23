package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	//*c += ByteCounter(len(p)) // convert int to ByteCounter
	*c = *c + ByteCounter(len(p))    // complain mismatching type without conversion
	return len(p), nil
}

func main(){
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly"

	c = 0
	c.Write([]byte("hello Dolly"))
	fmt.Println(c)
}
