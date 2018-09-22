package main

import (
	"fmt"
	"sync"
)


// don't know what this is for
type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	//ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- value
			println("index:", elem, ",value:", value)
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}

func main() {

	th := threadSafeSet{
		s: []interface{}{"a", "b"},
	}
	v := <-th.Iter()
	fmt.Printf("\noutgoing value:%v ",  v)
	//{
	//	v := <-th.Iter()
	//fmt.Printf("\noutgoing value:%v ",  v)
	//}
}
