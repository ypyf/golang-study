package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
	go func() {
		fmt.Println("Hello Go!")
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	fmt.Println("Program exit.")
}
