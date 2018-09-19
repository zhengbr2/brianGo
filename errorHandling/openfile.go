package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	//	fileNotExist()
	// pathError()
	safeFileNotExist()
	dnsError()
	glob()
	ignored()
}

func fileNotExist() {
	f, _ := os.Open("/test.txt")

	//if f == nil {
	//	fmt.Println("point f is nil")
	//	return
	//}
	fmt.Println(f.Name(), "opened successfully")
}

func safeFileNotExist() {
	isDone := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("handing panic here!", err)
			}
			isDone <- 1
		}()
		fileNotExist()
	}()
	<-isDone
}

func pathError() {
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Println("File at path", err.Path, "failed to open")
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

func dnsError() {
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println("temporary error")
		} else {
			fmt.Println("generic error: ", err)
		}
		return
	}
	fmt.Println(addr)
}

func glob() {
	files, error := filepath.Glob("[")
	if error != nil && error == filepath.ErrBadPattern {
		fmt.Println(error)
		return
	}
	fmt.Println("matched files", files)
}

func ignored() {
	files, _ := filepath.Glob("[")
	fmt.Println("matched files", files)
}
