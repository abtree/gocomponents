package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func new_day() {
	fmt.Printf("time.AfterFunc %d \n", GetGoid())
	time.AfterFunc(2*time.Second, new_day)
}

func main() {
	time.AfterFunc(2*time.Second, new_day)
	for {
		<-time.After(30 * time.Second)
		fmt.Printf("Finish %v \n", GetGoid())
		break
	}
}

// 这个函数可以变相获取goroutine的id 但是效率较差 所有一般只用于测试
func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine iF: %v", err))
	}

	return int64(id)
}
