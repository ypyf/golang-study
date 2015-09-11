package main

import "fmt"

func main() {
	ch := make(chan int32, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
