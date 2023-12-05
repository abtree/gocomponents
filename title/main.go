package main

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func main() {
	str := "yy_1_aaaa"
	fmt.Println(generator.CamelCase(str))
}
