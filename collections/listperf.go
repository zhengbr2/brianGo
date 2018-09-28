package main

import (
	"container/list"
	"time"
		"log"
)

func main() {
	sli:=make([]int ,10)
	before := time.Now()
	for i := 0; i<1*100000*1000;i++  {
		sli=append(sli, 1)
	}
	log.Println("appending slice::" + time.Now().Sub(before).String())

	before = time.Now()
	l:=list.New()
	for i := 0; i<1*100000*1000;i++  {
		l.PushBack(1)
	}
	log.Println("appending list::" + time.Now().Sub(before).String())

	// 比较遍历
	before = time.Now()
	for _,_ = range sli {
		//fmt.Printf("values[%d]=%d\n", i, item)
	}
	log.Println("walking through slice::" + time.Now().Sub(before).String())
	before = time.Now()
	for e := l.Front(); e != nil; e = e.Next() {
		//fmt.Println(e.Value)
	}
	log.Println("walking through list::" + time.Now().Sub(before).String())
}
