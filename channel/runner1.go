package main

import (
	"brianGo/channel/common"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := common.New(timeout)

	r.Add(createTask2(), createTask2(), createTask2())

	if err := r.Start(); err != nil {
		switch err {
		case common.ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case common.ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("...任务执行结束...")
}

func createTask2() func(int) {
	return func(id int) {
		log.Printf("正在执行任务%d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
