package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	w := sync.WaitGroup{}
	w.Add(2)

	v := []rune("Go Programming Language")

	go func() {
		defer w.Done()

		v[0] = 'X'
		time.Sleep(1 * time.Second)
		v[1] = 'Z'
	}()

	go func() {
		defer w.Done()
		fmt.Println(string(v)) // v = ?
	}()

	w.Wait()
	// v[0] = 'S'
	// time.Sleep(1 * time.Second)
	// v[1] = 'B'

	// fmt.Println(string(v)) // v = ?
}
