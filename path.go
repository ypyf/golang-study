package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	p, err := url.Parse("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	p.Path = path.Join(p.Path, "news")
	params := url.Values{}
	params.Set("appName", "ssss")
	p.RawQuery = params.Encode()
	fmt.Println(p.String())
}
