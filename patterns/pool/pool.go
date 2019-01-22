package pool

import "fmt"

type Object struct{}

func (o *Object ) Do(){
	fmt.Println("doing!")
}

type Pool chan *Object

func New(total int) Pool {
	p := make(Pool, total)
	for i := 0; i < total; i++ {
		p <- new(Object)
	}
	return p
}
