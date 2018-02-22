package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/net/proxy"
)

const (
	PROXY_ADDR = "ali.lerry.me:52344"
)

func newSocks5Proxy(addr string) proxy.Dialer {
	// create a socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", addr, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}
	return dialer
}

func ProxyServer(dialer proxy.Dialer, w http.ResponseWriter, r *http.Request) {

	log.Printf("original RequestURI: %s\n", r.RequestURI)

	// 设置HTTP客户端
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	uri := r.RequestURI
	if r.URL.Port() == "443" {
		uri = "https://" + r.RequestURI
	}
	log.Printf("New URI: %s\n", uri)
	// // 创建HTTP请求
	req, _ := http.NewRequest(r.Method, uri, r.Body)
	println(req.URL.Scheme)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println(err)
		w.Write([]byte("代理服务器请求失败"))
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	for _, value := range resp.Request.Cookies() {
		w.Header().Add(value.Name, value.Value)
	}

	w.WriteHeader(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("body :", err)
	}
	w.Write(body)
	log.Printf("完成代理服务: %v\n", req.URL)
}

func main() {
	dailer := newSocks5Proxy(PROXY_ADDR)
	r := mux.NewRouter()
	r.HandleFunc(`*/*`, func(w http.ResponseWriter, r *http.Request) { ProxyServer(dailer, w, r) })
	err := http.ListenAndServe("127.0.0.1:8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { ProxyServer(dailer, w, r) }))
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
