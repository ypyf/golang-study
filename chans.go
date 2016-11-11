package main

import (
	"fmt"
	"time"
)

func f(ccc chan chan chan int) {
	fmt.Printf("f(%#v)\n", ccc)
	cc := <-ccc
	go g(cc)
	cc <- make(chan int)
}

func g(cc chan chan int) {
	fmt.Printf("g(%#v)\n", cc)
	c := <-cc
	go h(c)
	c <- 1
}

func h(c chan int) {
	fmt.Printf("h(%#v)\n", c)
}

func main() {
	a := make(chan chan chan int)
	go f(a)
	a <- make(chan chan int)
	time.Sleep(1 * time.Second)
}
