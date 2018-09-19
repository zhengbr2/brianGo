package main

func main() {
	ch := make(chan int)
	ch <- 5
	a := <-ch
	print(a)
}
