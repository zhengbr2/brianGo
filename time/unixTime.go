package main

import "time"

func main() {

	ellapse:=time.Now().Unix()
	println(ellapse)
	diffSec:=(ellapse -1541574160)
	println(diffSec)

}
