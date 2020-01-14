package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	urls := []string{
		"http://stackoverflow.com",
		"http://www.douban.com",
		"http://www.zhihu.com",
		"http://www.google.com", // 模拟故障
		"http://www.baidu.com",
		"http://www.bilibili.com"}
	var g errgroup.Group
	for _, url := range urls {
		url := url
		g.Go(func() error {
			client := http.Client{
				Timeout: 5 * time.Second,
			}
			resp, err := client.Get(url)
			if err == nil {
				server := resp.Header.Get("Server")
				if server == "" {
					server = "Unknown"
				}
				fmt.Printf("%s Server: %s\n", url, server)
				resp.Body.Close()
			}
			return err
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println("Successfully fetched all hosts.")
	}
}
