package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func main() {
	const ADDR = "127.0.0.1:5555"
	s, err := net.Listen("tcp", ADDR)
	checkError(err)
	log.Printf("AI listening on %s\n", s.Addr())

	for {
		c, err := s.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func() {
			reader := bufio.NewReader(c)
			for {
				s, err := reader.ReadString('\n')
				checkError(err)
				log.Println(s)
				s = strings.Replace(s, "吗", "", 1)
				s = strings.Replace(s, "?", "!", 1)
				_, err = io.Copy(c, strings.NewReader(s))
				checkError(err)
			}
		}()
	}
}
