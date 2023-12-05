package main

import (
	"fmt"
	"io"
	"net/http"
)

type WebTest struct {
}

func (*WebTest) SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (*WebTest) SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}

var mux *http.ServeMux

func init() {
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello Tls")
	})
}

func main() {
	//带证书的https访问
	go http.ListenAndServeTLS(":4430", "../grpc/certs/server.pem", "../grpc/certs/server.key", mux)
	var w WebTest
	http.Handle("/", http.HandlerFunc(w.SayHello))
	http.Handle("/hi", http.HandlerFunc(w.SayHi))
	//普通的http服务
	http.ListenAndServe("localhost:8090", nil)
}
