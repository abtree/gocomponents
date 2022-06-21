package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("txt.go", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Open File Error")
	}
	var builder strings.Builder
	var enter string = "\r\n"
	file.WriteString("package main" + enter)
	file.WriteString(enter)
	file.WriteString(`import "fmt"` + enter)
	file.WriteString(enter)
	file.WriteString("func Print(){" + enter)
	file.WriteString(`fmt.Println("Hallo World")` + enter)
	file.WriteString("}" + enter)

	builder.WriteString("func ToString() string{" + enter)
	builder.WriteString(`return ""` + enter)
	builder.WriteString("}" + enter)
	file.WriteString(builder.String())
}
