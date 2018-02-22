package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/proxy"
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

func handleClientRequest(dialer proxy.Dialer, client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	buf := make([]byte, 4096)
	n, err := client.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}

	sep := bytes.IndexByte(buf, '\n')
	if sep == -1 {
		log.Printf("Bad HTTP Request\n")
		return
	}

	log.Printf("代理请求: %s\n", buf[:sep])

	// 解析URL
	var method, host, address string
	fmt.Sscanf(string(buf[:sep]), "%s%s", &method, &host)
	if method == "CONNECT" {
		address = host
	} else {
		hostPortURL, err := url.Parse(host)
		if err != nil {
			log.Println(err)
			return
		}
		if hostPortURL.Opaque == "443" {
			address = hostPortURL.Scheme + ":443"
		} else {
			if strings.Index(hostPortURL.Host, ":") == -1 {
				address = hostPortURL.Host + ":80"
			} else {
				address = hostPortURL.Host
			}
		}
	}

	var server net.Conn

	if dialer != nil {
		server, err = dialer.Dial("tcp", address)
	} else {
		server, err = net.Dial("tcp", address)
	}

	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(buf[:n])
	}

	go io.Copy(server, client)
	io.Copy(client, server)
}

func startSocks5Proxy(useProxy bool) {
	var dailer proxy.Dialer = nil
	if useProxy {
		dailer = newSocks5Proxy(os.Getenv("SOCKS5_PROXY_ADDR"))
	}
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(dailer, client)
	}
}

func startHttp5Proxy() {
	println("HTTP Proxy 未实现")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "socks5":
			startSocks5Proxy(true)
		case "http":
			startHttp5Proxy()
		default:
			println("usage: socksproxy [socks5 | http]")
		}
	} else {
		startSocks5Proxy(false)
	}
}
