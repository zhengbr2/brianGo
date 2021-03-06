package main

import (
	"fmt"
	"time"
)

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
	fmt.Println("squareop <- sum")
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
	fmt.Println("cubeop <- sum")
}

func main() {

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Print(squares, " ", cubes, "\n")
	fmt.Println("Final output", squares+cubes)

	time.Sleep(time.Millisecond * 100)
	ci := make(chan int)

	go func(c chan int) {
		c <- 1
	}(ci)

	ri, ok := <-ci
	println(ri)
	println(ok)
	//ra:=<-ci  //blockblock

	//inta,ok:=<-ci   //lobck

	close(ci) // must close , other wise deadlock

	ia, ok := <-ci
	fmt.Print("ra read:", ia, ok)

}
