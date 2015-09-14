package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.ListenAndServe(":8080", r)
}
