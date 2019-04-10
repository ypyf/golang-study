package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	w := sync.WaitGroup{}
	w.Add(2)

	v := 0

	go func() {
		defer w.Done()
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)
		v++
	}()

	time.Sleep(time.Duration(rand.Int31n(3)) * time.Second)

	go func() {
		defer w.Done()
		v--
	}()

	fmt.Println(v)

	w.Wait()
}
