package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"strconv"
	"time"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var now=time.Now()
	wg.Add(200)
	for j := 0; j < 200; j++ {


		go gethttp(j)
	}
	wg.Wait()
	pass := time.Now().Sub(now).Seconds()
	log.Println("elapse:", pass)
}

func gethttp(j int) {

	for i := 0; i < 100; i++ {
		resp, err := http.Get("http://localhost:9090")

		if err != nil {

			//log.Fatal(err)
			continue

		}

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {

			// handle error
			continue

		}

		fmt.Println(string(body) + strconv.Itoa(j*100+i))
		//resp.Body.Close()


	}
	wg.Done()

}
