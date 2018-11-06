package main

import (
	"brianGo/serialization/proto"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/proto"
	"log"
	"strconv"
	"time"
)

type User struct {
	Age    int32
	Name   string
	BornAt string
}

func main() {
	looptimes := 10000000
	u := User{66, "nxin", "beijing"}
	gobbegintimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	gobbeginint, _ := strconv.Atoi(gobbegintimestamp)
	fmt.Println("gob序列化==============================", gobbeginint)
	buf := new(bytes.Buffer)   //分配内存
	enc := gob.NewEncoder(buf) //创建基于buf内存的编码器
	for i := 0; i < looptimes; i++ {

		err := enc.Encode(u) //编码器对结构体编码
		if err != nil {
			log.Fatal(err)
		}
	}
	gobendtimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	gobendint, _ := strconv.Atoi(gobendtimestamp)
	fmt.Println("===================END===================", gobendint)

	jsonbegintimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	jsonbeginint, _ := strconv.Atoi(jsonbegintimestamp)
	fmt.Println("json序列化==============================", jsonbeginint)
	for j := 0; j < looptimes; j++ {
		_, e := json.Marshal(u)
		if e != nil {
			log.Fatal(e)
		}
	}
	jsonendtimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	jsonendint, _ := strconv.Atoi(jsonendtimestamp)
	fmt.Println("===================END===================", jsonendint)

	protobufbegintimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	protobufbeginint, _ := strconv.Atoi(protobufbegintimestamp)
	fmt.Println("protobuf序列化==============================", protobufbeginint)
	hw := &myproto.User{Age: 10, Name: "wang", BornAt: "beijing"}
	for j := 0; j < looptimes; j++ {
		_, e := proto.Marshal(hw)
		if e != nil {
			log.Fatal(e)
		}
	}
	protobufendtimestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	protobufendint, _ := strconv.Atoi(protobufendtimestamp)
	fmt.Println("===================END===================", protobufendint)

	fmt.Println("json:", time.Duration(jsonendint-jsonbeginint).Seconds())
	fmt.Println("gob:", time.Duration(gobendint-gobbeginint).Seconds())
	fmt.Println("protobuf:", time.Duration(protobufendint-protobufbeginint).Seconds())

}
