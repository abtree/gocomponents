package main

import (
	"fmt"
)

type IBaser1 interface {
	Show()
}

type IBaser2 interface {
	Print()
}

type Test struct {
}

func (*Test) Show() {
	fmt.Println("IBaser1. Show")
}

func (*Test) Print() {
	fmt.Println("IBaser2. Print")
}

func Brival1(impl IBaser1) {
	impl.Show()
	Brival2(impl.(IBaser2)) //此处将接口1转换为接口2
}

func Brival2(impl IBaser2) {
	impl.Print()
}

func main() {
	t := &Test{} //t同时实现了两个接口
	Brival1(t)
}
