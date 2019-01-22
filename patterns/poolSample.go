package main

import "brianGo/patterns/pool"

func main() {
	p := pool.New(2)
	select {
	case obj := <-p:
		obj.Do()
		p <- obj
	default:
		// No more objects left — retry later or fail
		return
	}
}
