// http代理(Proxy Server) project main.go
// 1871522910@qq.com
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(r.Method, r.RequestURI, r.Body)

	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	req.Host = r.Host
	req.URL.Scheme = r.URL.Scheme
	req.URL.Host = r.URL.Host
	req.URL.Path = r.URL.Path
	req.Proto = r.Proto
	req.ProtoMajor = r.ProtoMajor
	req.ProtoMinor = r.ProtoMinor
	req.Close = r.Close

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
	log.Printf("完成代理服务: %v\n", r.URL)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
