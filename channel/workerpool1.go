package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"net/http"
	"runtime/pprof"
)

// playload
type Payload struct {
	name string
}

func (p *Payload) Handle() {
	fmt.Printf("%s Handling tasking ...\n", p.name)
}

type Job struct {
	Payload Payload
}

//job queue
var JobQueue chan Job
var wg = &sync.WaitGroup{}

type Worker struct {
	name       string
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job, name string) Worker {
	fmt.Printf("创建了一个工人,它的名字是:%s \n", name)
	return Worker{
		name:       name,
		WorkerPool: workerPool,
		//default value
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			//注册到对象池中,
			w.WorkerPool <- w.JobChannel
			//fmt.Printf("[%s]把自己注册到 对象池中 \n", w.name)
			select {
			//接收到了新的任务
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任务 当前空闲工人数是[%d]\n", w.name, len(w.WorkerPool))
				job.Payload.Handle()
				wg.Done()
				time.Sleep(time.Millisecond * 1000)
				//接收到了任务
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	//WorkerPool chan JobQueue
	name       string
	maxWorkers int
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		name:       "dispacher",
		maxWorkers: maxWorkers,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		worker.Start()
	}
	go d.dispatch()
}

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}


func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			fmt.Println("调度者,接收到一个工作任务")
			//time.Sleep(10* time.Millisecond)
			// 调度者接收到一个工作任务
		go	func(job Job) {    //这里开启了无限协程
				//从现有的对象池中拿出一个
				jobChannel := <-d.WorkerPool
				jobChannel <- job

			}(job)
		}
	}
}

func initialize() {
	maxWorkers := 8000 //池子大小
	maxQueue := 10      //指定任务的队列长度
	//初始化一个调度者,并指定它可以操作的 工人个数
	dispatcher := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue)

	dispatcher.Run()
}

func main() {


go func() {
	http.HandleFunc("/dump", dumpHandler)
	http.ListenAndServe(":8080", nil)
}()



	before := time.Now()
	initialize()
	wg.Add(10000000)
	for i := 0; i < 10000000; i++ {
		p := Payload{
			fmt.Sprintf("playload-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Payload: p,
		}
		//time.Sleep(time.Millisecond * 1)
	}
	fmt.Println("任务派遣完毕", time.Now())
	//time.Sleep(time.Second * 1)
	wg.Wait()
	close(JobQueue)
	fmt.Println("耗时：", time.Now().Sub(before).Seconds())
}
