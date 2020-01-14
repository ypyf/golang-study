package main

import (
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	go func() { log.Fatal(http.ListenAndServe(":http", nil)) }()
	select {
	}
}
