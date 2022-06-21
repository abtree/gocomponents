package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type SPaths struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type SStructs struct {
	Name   string `json:name`
	Field  string `json:field`
	Struct string `json:struct`
}

type Data struct {
	Paths   []*SPaths   `json:paths`
	Package []string    `json:"packages"`
	Structs []*SStructs `json:structs`
}

func main() {
	data := &Data{
		Paths:   make([]*SPaths, 0),
		Package: make([]string, 0),
		Structs: make([]*SStructs, 0),
	}
	byts, err := ioutil.ReadFile("build.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(byts, data)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v", data)
}
