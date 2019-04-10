package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	go func(x ...int) {
		fmt.Printf("%v\n", x)
	}(1, 2, 3)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	<-sig
	fmt.Println("Program exit (Ctrl+C).")
}
