package main

import (
	"fmt"
	"time"
)

func broadcast(msg string, delay time.Duration) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay * time.Second)
		fmt.Printf("New message: %s\n", msg)
		close(ch)
	}()
	return ch
}

func main() {
	wait := broadcast("hello world", 3)
	fmt.Println("Waiting for the message...")
	<-wait
	fmt.Println("The message is out, time to leave.")
}
