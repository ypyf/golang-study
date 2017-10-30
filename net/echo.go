package main

import (
	"io"
	"log"
	"net"
)

func main() {
	s, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		c, err := s.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go io.Copy(c, c)
	}
}
