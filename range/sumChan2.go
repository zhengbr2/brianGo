package main

import (
		"fmt"

	"sync"
	"time"
)

func sum(values []int, resultChan chan int, wg * sync.WaitGroup) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	resultChan <- sum // 将计算结果发送到channel中

	wg.Done()
}
func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,11,12,13,14,15,16}
	workerCount := 4
	resultChan := make(chan int, workerCount)
	var wg sync.WaitGroup
	wg.Add(workerCount)
	for i:=0; i < workerCount; i++ {
		go sum (values[i*16/workerCount: (i+1)*16/workerCount], resultChan, &wg)
	}

	wg.Wait()
	close(resultChan)
	total:=0

	for i:= range resultChan{
		fmt.Println("read:", i)
		total = i+ total
	}
	println(total)
}
