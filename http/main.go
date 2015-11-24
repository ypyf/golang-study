package main

import (
	"log"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("www.baidu.com")
	if err != nil {
		log.Fatalf("error %v\n", err)
	}

	io.Copy(os.Stdout, resp.Body)
}
