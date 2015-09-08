package main

import "fmt"

func job(ch chan int, i int) {
	ch <- i
}

func main() {

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go job(ch, i)
		fmt.Println(<-ch)
	}

}
