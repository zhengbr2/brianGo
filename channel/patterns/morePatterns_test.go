package patterns

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func or(chans ...<-chan interface{}) <-chan interface{} {
	tmp := make(chan struct{})
	var a interface{}
	out := make(chan interface{})
	go func() {
		var once sync.Once
		for _, c := range chans {
			go func(c <-chan interface{}) {
				select {
				case a = <-c:
					once.Do(func() { close(tmp) })
					go func() {
						out <- a
					}()
				case <-tmp:
				}
			}(c)
		}
	}()
	return out
}


func TestOr(t *testing.T) {
	c1, c2, c3 := make(chan interface{}), make(chan interface{}), make(chan interface{})
	out := or(c1, c2, c3)

	go func() {
		time.Sleep(time.Millisecond * 1)
		c1 <- "C1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c3 <- "C3"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c2 <- struct {
			string
			int
			bool
		}{"C2", 100, true}
	}()
	fmt.Printf("what you input is:%v", <-out)
}



func or2(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		//reflect.Select(cases)
		chosen, recv, ok := reflect.Select(cases)
		if ok {
			fmt.Println("chosen:",chosen)
			fmt.Println("value:",recv.Interface())
			go func() {
				orDone<-recv.Interface()
			}()
		}
	}()
	return orDone
}

func TestOr2(t *testing.T) {
	c1, c2, c3 := make(chan interface{}), make(chan interface{}), make(chan interface{})
	out := or2(c1, c2, c3)

	go func() {
		time.Sleep(time.Millisecond * 1)
		c1 <- "C1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c3 <- "C3"
	}()

	go func() {
		//time.Sleep(time.Millisecond * 1)
		c2 <- struct {
			string
			int
			bool
		}{"C2", 100, true}
	}()
	fmt.Printf("what you input is:%v", <-out)
}


func or3(channels ...<-chan interface{}) <-chan interface{} {
	// not work yet

	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	tmp := make(chan struct{})
	done := make(chan interface{})
	go func() {
		defer close(tmp)
		switch len(channels) {
		case 2:
			select {
			case v:= <-channels[0]:
				go func() {done<-v}()
			case v:=<-channels[1]:
				go func() {done<-v}()
			}
		default:
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return done
}

func TestOr3(t *testing.T) {
	c1, c2, c3 := make(chan interface{}), make(chan interface{}), make(chan interface{})
	out := or3(c1)

	go func() {
		time.Sleep(time.Millisecond * 1200)
		c1 <- "C1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c3 <- "C3"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c2 <- struct {
			string
			int
			bool
		}{"C2", 100, true}
	}()
	time.Sleep(100)
	fmt.Printf("what you input is:%v", <-out)
}


func orDone(done <-chan struct{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if ok == false {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}


func fanIn(chans ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))
		for _, c := range chans {
			go func(c <-chan interface{}) {
				for v := range c {
					fmt.Printf("received %v\n",v)
					out <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}


func TestFanin(t *testing.T) {
	c1, c2, c3 := make(chan interface{}), make(chan interface{}), make(chan interface{})
	out := fanIn(c1,c2,c3)

	go func() {
		time.Sleep(time.Millisecond * 1)
		c1 <- "C1"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c3 <- "C3"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)
		c2 <- struct {
			string
			int
			bool
		}{"C2", 100, true}
	}()
	//time.Sleep(time.Second*1)
	fmt.Printf("what you input is:%v", <-out)
}