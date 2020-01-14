package main

import (
	"expvar"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("客户端请求: %v %v\n", r.RemoteAddr, r.RequestURI)
	w.Write([]byte("响应：\n"))
	fmt.Fprintf(w, "方法：%v\n", r.Method)
	fmt.Fprintf(w, "URL：%v\n", r.URL)
	fmt.Fprintf(w, "协议：%v\n", r.Proto)
	fmt.Fprintf(w, "主机：%v\n", r.Host)
	fmt.Fprintf(w, "内容长度：%v\n", r.ContentLength)
	for k := range r.Header {
		fmt.Fprintf(w, "%v：%v\n", k, strings.Join(r.Header[k], ", "))
	}
}

func goroutines() interface{} {
	return runtime.NumGoroutine()
}

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/{articles:[a-zA-Z]+}", handler)
	go http.ListenAndServe(":8080", r)

	// 监控
	expvar.Publish("Goroutines", expvar.Func(goroutines))
	go http.ListenAndServe(":1234", nil)
	ch := make(chan struct{})
	<-ch
}
