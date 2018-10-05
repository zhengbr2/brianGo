package main

import (
	"log"
	"sync"
)

type T int

func IsClosed(ch <-chan T) bool {
	select {
	case a := <-ch:
		log.Printf("value of a:%d", a)
		return true
	default:
		return false
	}

}

func SafeSend(ch chan T, value T) (closed bool) {

	defer func() {
		if recover() != nil {
			// the return result can be altered
			// in a defer function call
			closed = true
		}
	}()

	ch <- value  // panic if ch is closed
	return false // <=> closed = false; return
}

func SafeClose(ch chan T) (justClosed bool) {
	defer func() {
		if recover() != nil {
			justClosed = false
		}
	}()

	// assume ch != nil here.
	close(ch) // panic if ch is closed
	return true
}

type MyChannel struct {
	C    chan T
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan T)}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func main() {

	c := make(chan T)
	log.Println(IsClosed(c)) // false
	close(c)
	log.Println(IsClosed(c)) // true

	ret := SafeSend(c, 1)
	log.Println("channel closed:", ret)

	ret = SafeClose(c)
	log.Println("channel is newly closed:", ret)

	mc := NewMyChannel()
	mc.SafeClose()
	mc.SafeClose()

	mc2 := NewMyChannel2()
	log.Println("mc2 closed:", mc2.IsClosed())
	mc2.SafeClose()
	log.Println("mc2 closed:", mc2.IsClosed())

}

type MyChannel2 struct {
	C      chan T
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel2() *MyChannel2 {
	return &MyChannel2{C: make(chan T)}
}

func (mc *MyChannel2) SafeClose() {
	mc.mutex.Lock()
	if !mc.closed {
		close(mc.C)
		mc.closed = true
	}
	mc.mutex.Unlock()
}

func (mc *MyChannel2) IsClosed() bool {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	return mc.closed
}
