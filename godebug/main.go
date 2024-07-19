package main

import (
	"fmt"
	"time"
)

/*
	windows 上使用 环境变量 GODEBUG (一个强大的go调试信息环境变量，有非常多的调试命令)

需要先运行命令 (set 设置环境变量)

	: set GODEBUG=schedtrace=1000
	: set GODEBUG=allocfreetrace=1
	: set GODEBUG=clobberfree=1
	: set GODEBUG=gctrade=1

再执行go应用程序，即可看到调试信息
*/
func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello debug %d \n", i)
	}
}
