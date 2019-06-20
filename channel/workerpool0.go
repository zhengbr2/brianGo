package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobQue = make(chan Job, 8)
var resultQue = make(chan Result, 3)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	//time.Sleep(1 * time.Second)
	return sum
}
func worker(wg *sync.WaitGroup, id int) {
	for job := range jobQue {
		output := Result{job, digits(job.randomno)}
		resultQue <- output
		fmt.Printf("have just handle job id %d, job.randomno %d by worker# %d\n", job.id, job.randomno, id)
		time.Sleep(time.Millisecond * 200)
	}
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg, i)
	}
	wg.Wait()
	close(resultQue)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobQue <- job
	}
	close(jobQue)
}
func result(done chan bool) {
	for result := range resultQue {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	fmt.Printf("now all result are done!\n")
	done <- true
}
func main() {
	startTime := time.Now()
	noOfJobs := 1000000
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 40000
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
