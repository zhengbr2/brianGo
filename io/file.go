package main

import (
	"log"
	"os"
)

func main() {
	os.Mkdir("brianzheng", 0777)
	os.MkdirAll("brianzheng/test1/test2", 0777)
	err := os.Remove("brianzheng")
	if err != nil {
		log.Println(err)
	}
	err = os.RemoveAll("brianzheng")
	if err != nil {
		log.Println(err)
	}
	log.Println("all folders removed")

	fname := "zhengbr2.txt"
	err = os.Remove("zhengbr2.txt")
	if err != nil {
		log.Printf("not able to remove file %v", err)
	}
	fout, err := os.Create(fname)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		fout.WriteString("just a test\r\n")
		fout.Write([]byte("just a test\r\n"))
	}
	fout.Close()

	fl, err := os.Open(fname)
	defer fl.Close()

	if err != nil {
		log.Println(err)
		return
	}
	var buf []byte
	buf = make([]byte, 16)
	//buf = [1024]byte{}[:]

	for {
		n, _ := fl.Read(buf)
		if n == 0 {
			break
		}
		print(string(buf[0:n]))
	}

}
