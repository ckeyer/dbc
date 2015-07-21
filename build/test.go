package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func SayHi(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hi, this is a test app"))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/hi", SayHi)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(err.Error())
	}
}
