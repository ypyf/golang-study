package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var result int
	processors := runtime.GOMAXPROCS(0)
	for i := 0; i < processors; i++ {
		go func() {
			for {
				result++
			}
		}()
	}
	time.Sleep(time.Second) //wait for go function to increment the value.
	fmt.Println("result =", result)
}
