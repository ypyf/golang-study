package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 99; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Println(num)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(100)
}
