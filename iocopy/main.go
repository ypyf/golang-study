package main

import (
    "bytes"
    "fmt"
    "io"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "baidu.com:80")
    if err != nil {
        fmt.Println("dial error:", err)
        return
    }
    defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n")
	fmt.Fprintf(conn, "Host: www.baidu.com\r\n")
	fmt.Fprintf(conn, "Connection: close\r\n")
	fmt.Fprintf(conn, "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36\r\n")
	fmt.Fprintf(conn, "Pragma: no-cache\r\n")
	fmt.Fprintf(conn, "Cache-Control: no-cache\r\n")
	fmt.Fprintf(conn, "\r\n")
	var buf bytes.Buffer
    io.Copy(&buf, conn)
    fmt.Printf("%s\n", string(buf.Bytes()))
    fmt.Println("total size:", buf.Len())
}
