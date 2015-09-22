// http代理(Proxy Server) project main.go
// 1871522910@qq.com
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	//log.Println(r.RequestURI)
	req, _ := http.NewRequest(r.Method, r.RequestURI, r.Body)
	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	targetURL, _ := url.Parse("http://www.ifeng.com")
	req.Host = targetURL.Host
	req.URL.Scheme = targetURL.Scheme
	req.URL.Host = targetURL.Host
	req.URL.Path = targetURL.Path
	req.Proto = "HTTP/1.1"
	req.ProtoMajor = 1
	req.ProtoMinor = 1
	req.Close = false

	resp, err := http.DefaultClient.Do(req)
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
	log.Println("成功完成一次代理服务;url:", r.URL)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
