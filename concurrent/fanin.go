package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2, input3 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 30; i++ {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case s := <-input3:
				c <- s
			}
		}
		close(c)
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Mark"), boring("Ann"))
	for x := range c {
		fmt.Println(x)
	}
	fmt.Println("You're boring; I'm leaving.")
}
