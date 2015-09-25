package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func ProxyServer(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(r.Method, r.RequestURI, r.Body)
	log.Println(r.URL)

	for k, v := range r.Header {
		for _, vv := range v {
			req.Header.Add(k, vv)
		}
	}
	req.Host = "192.168.4.4"
	req.URL.Scheme = "http"
	req.URL.Host = req.Host
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
	r := mux.NewRouter()
	r.HandleFunc(`/doc/{doc}`, ProxyServer)
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
