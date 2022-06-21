package main

import (
	"fmt"
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

func main() {
	var w WebTest
	http.Handle("/", http.HandlerFunc(w.SayHello))
	http.Handle("/hi", http.HandlerFunc(w.SayHi))
	http.ListenAndServe("localhost:8090", nil)
}
