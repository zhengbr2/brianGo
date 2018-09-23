package main

import "log"

func main() {

	log.SetFlags(log.Ldate|log.Lmicroseconds|log.Lshortfile)
	log.Println("brian is testing logger")
	log.Println("end of log")
	log.Fatalf("fatal and quit!!!")
	log.Println("not expected here")

}