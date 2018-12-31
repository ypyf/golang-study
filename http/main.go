package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalf("error %v\n", err)
	}

	io.Copy(os.Stdout, resp.Body)
}
