package main

import (
	"components/websocket/controllers"
	"fmt"
	"log"
	"net/http"
)

type HandleFnc func(http.ResponseWriter, *http.Request)

//处理异常的闭包封装函数
func logPanics(f HandleFnc) HandleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		w.Header().Add("Access-Control-Max-Age", "3600")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,Access-Token,Authorization,ybg")
		f(w, r)
	}
}

func init() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("views"))))
	http.HandleFunc("/ws/echo", logPanics(controllers.BaseController.Init))
}

func main() {
	fmt.Println("server start：8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: 8081", err)
	}
}
