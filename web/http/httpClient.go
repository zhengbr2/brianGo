package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	var now = time.Now()
	wg.Add(20)
	for j := 0; j < 20; j++ {

		go gethttp(j)
	}
	wg.Wait()
	pass := time.Now().Sub(now).Seconds()
	log.Println("elapse:", pass)
}

func gethttp(j int) {

	for i := 0; i < 1000; i++ {
		resp, err := http.Get("http://localhost:9090")
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body) + strconv.Itoa(j*100+i))
		//resp.Body.Close()
	}
	wg.Done()

}
