package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	go func(x ...int) {
		fmt.Printf("%v\n", x)
	}(1, 2, 3)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		<-sig
		fmt.Println("Program exit (Ctrl+C).")
		wg.Done()
	}()

	wg.Wait()
}
