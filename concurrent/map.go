package main

import (
	"fmt"
	"sync"
	"time"
)

var a map[string]int
var wg sync.WaitGroup

func main() {
	a = make(map[string]int)
	a["aaa"] = 1
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(850 * time.Millisecond)
			if _, ok := a["aaa"]; !ok {
				panic("错误")
			}
			fmt.Println("Ok")
		}()
	}
	wg.Wait()
}
