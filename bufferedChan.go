package main

import "fmt"

func main() {
	ch := make(chan int32, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	ch <- 10
	close(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
