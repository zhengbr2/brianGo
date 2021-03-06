package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

var ErrTimeOut = errors.New("Timeout!")
var ErrInterrupt = errors.New("Interrupted")
var ErrApplication = errors.New("Application Encounters an Error!")

//一个执行者，可以执行任何任务，但是这些任务是限制完成的，
//该执行者可以通过发送终止信号终止它
type Runner struct {
	tasks         []func(int) error //要执行的任务
	runningStatus chan error        //用于通知任务全部完成
	timeout       <-chan time.Time  //这些任务在多久内完成
	interrupt     chan os.Signal    //可以控制强制终止的信号
}

func New(tm time.Duration) *Runner {
	return &Runner{
		runningStatus: make(chan error),
		timeout:       time.After(tm),
		interrupt:     make(chan os.Signal, 1),
	}
}

//将需要执行的任务，添加到Runner里
func (r *Runner) Add(tasks ...func(int) error) {
	r.tasks = append(r.tasks, tasks...)
}

//执行任务，执行的过程中接收到中断信号时，返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) runOnSequence() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		e := task(id + 1)
		if e != nil {
			return e
		}
	}
	return nil
}

//检查是否接收到了中断信号
func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *Runner) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.runningStatus <- r.runOnSequence()
	}()

	select {
	case err := <-r.runningStatus:
		if err == ErrApplication {
			return ErrApplication
		} else {
			return err
		}
	case <-r.timeout:
		return ErrTimeOut
	}
}

func main() {
	log.Println("...开始执行任务...")

	timeout := 4 * time.Second
	r := New(timeout)
	time.Sleep(20 * time.Microsecond)

	r.Add(createTask2(), createTask2(), createTask2())
	err := r.Start()
	if err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		case ErrApplication:
			log.Println(err)
			os.Exit(3)

		}
	} else {
		log.Println("no error", err)
	}
	log.Println("...任务执行结束...")
}

func createTask2() func(int) error {
	return func(id int) error {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
		if id == 2 {
			return ErrApplication
		}
		log.Printf("执行完毕：%d", id)
		return nil
	}
}
