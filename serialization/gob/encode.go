package main

import (
	"encoding/gob"
	"log"
	"os"
	"fmt"
)

type Address struct {
	Type             string
	City             string
	Country          string
}

type VCard struct {
	FirstName    string
	LastName    string
	Addresses    []*Address
	Remark        string
}

var content    string

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	pa2 := &Address{"private", "Aartselaar", "Belgium"}
	wa2 := &Address{"work", "Boom", "Belgium"}
	pa3 := &Address{"private", "Aartselaar", "Belgium"}
	wa3 := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa,pa2, wa2,pa3, wa3}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// using an encoder:
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)

	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	file.Close()

	vc2 := VCard{}
	file2, _ := os.OpenFile("vcard.gob", os.O_RDONLY, 0666)
	dec:=gob.NewDecoder(file2)
	err=dec.Decode(&vc2)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	file2.Close()
	fmt.Println("%v",vc2)
	fmt.Printf("xxx",vc2.Addresses[0],vc2.Addresses[1])
}
