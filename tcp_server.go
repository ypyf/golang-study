package main

import (
	"bytes"
	"io"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":http")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go func() {
				var buf bytes.Buffer
				_, err := io.Copy(&buf, conn)
				if err != nil {
					log.Println(err)
				} else {
					log.Printf("%s\n", buf.String())
					conn.Write([]byte("Hello"))
					conn.Close()
				}
			}()
		}
	}
}
