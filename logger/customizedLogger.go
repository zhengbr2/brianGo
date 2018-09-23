package main

import (
	"log"
	"os"
	"io"
	"time"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)

func init(){
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("error open file：",err)
	}

	Info = log.New(os.Stdout,"Info:",log.Ldate | log.Lmicroseconds | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Lmicroseconds | log.Lshortfile)
	time.Sleep(10000)
	Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Lmicroseconds | log.Lshortfile)

}

func main() {
	Info.Println("running fine")
	Warning.Printf("Attention：%s\n","typhon is comming")
	Error.Println("Error msg throw here!")
}
