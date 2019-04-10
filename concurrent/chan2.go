package main

import "fmt"

func task(ch chan int) {
	for {
		i := <-ch
		fmt.Printf("task %v\n", i)
		i++
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go task(ch)
	ch <- 0
	for {
		i := <-ch
		if i > 10 {
			break
		}
		fmt.Printf("main %v\n", i)
		i++
		ch <- i
	}
}
