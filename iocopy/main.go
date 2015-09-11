package main

import (
    "fmt"
    "net"
    "io"
    "bytes"
)

func main() {
    conn, err := net.Dial("tcp", "baidu.com:80")
    if err != nil {
        fmt.Println("dial error:", err)
        return
    }
    defer conn.Close()
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    var buf bytes.Buffer
    io.Copy(&buf, conn)
    fmt.Printf("%s\n", string(buf.Bytes()))
    fmt.Println("total size:", buf.Len())
}