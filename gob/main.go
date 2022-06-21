package main

/*
gob 是go语言特有的序列化和反序列化方法
*/

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string

func main() {
	Read()
}

func Write() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println(err)
	}
}

func Read() {
	var vc VCard
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	dec := gob.NewDecoder(file)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println(err)
	}
	log.Println(vc)
}
