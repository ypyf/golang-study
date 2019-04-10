package main

import "fmt"
import "time"

func job(ch chan int, i int) {
	time.Sleep(1 * time.Second) // 睡眠1秒钟
	ch <- i
}

func main() {

	ch := make(chan int, 10)
	for i := 10; i > 0; i-- {
		go job(ch, i)
		fmt.Println(<-ch)
	}

}
