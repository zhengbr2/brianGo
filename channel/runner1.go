package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

//一个执行者，可以执行任何任务，但是这些任务是限制完成的，
//该执行者可以通过发送终止信号终止它
type Runner struct {
	tasks     []func(int)      //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //这些任务在多久内完成
	interrupt chan os.Signal   //可以控制强制终止的信号

}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

//将需要执行的任务，添加到Runner里
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行任务，执行的过程中接收到中断信号时，返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil    //nil 可以send到 error，但是不能send 到string
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
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

type  person struct{
	age int
	name string
}
func main() {
	log.Println("...开始执行任务...")

	{

		s:= make (chan error)
		go func () {s<-nil}()  //nil只能赋值给指针、channel、func、interface、map或slice类型的变量
		time.Sleep(111)
		select {
			case output:=<-s:
					println(output ==nil)
					//println("I have output here!" + output.Error())
			default:
				println("default output")
		}
		return
	}

	timeout := 5 * time.Second
	r := New(timeout)

	r.Add(createTask1(), createTask1(), createTask1())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask1() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
