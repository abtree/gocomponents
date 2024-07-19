package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

// runtime/trace调试
// 运行 go tool trace trace.out
// 即可在浏览器查看结果
func main() {
	// 1.创建trace文件
	file, err := os.Create("trace.out")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	// 2.开始trace
	trace.Start(file)
	//需要trace的内容(这里简单写个打印函数)
	fmt.Println("Hello Trace")
	// 3.结束trace
	trace.Stop()
}
