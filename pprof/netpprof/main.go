package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// 打开浏览器 输入网址 http://localhost:6060 即可查看pprof信息
func main() {
	//程序需要监控的逻辑
	fmt.Println("Hello pprof")
	//程序需要监控的逻辑
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
