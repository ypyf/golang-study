package main

import (
	"log"
	"math/rand"
	"time"
)

func Worker(sem chan int, lock chan<- struct{}, id int) {
	sem <- 1 // down(P原语)
	log.Println(id)
	time.Sleep(time.Duration(rand.Int31n(5)) * time.Second)
	<-sem // up(V原语)

	lock <- struct{}{}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan int, 2) // 带2个int缓冲的Channel(异步Channel，在读完或写满之前线程都不会阻塞)
	lock := make(chan struct{})
	for i := 0; i < 10; i++ {
		go Worker(ch, lock, i)
	}

	// 等待所有线程退出
	for i := 0; i < 10; i++ {
		<-lock
	}
	println("所有goroutine退出")
}
