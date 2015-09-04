package main

import "fmt"
import "time"

var count = 0

func Worker(sem chan int, lock chan bool, id int) {
    sem <- 1    // down(P原语)
    fmt.Println(time.Now().Format("15:04:05"), id)
    time.Sleep(1 * time.Second) // 睡眠1秒钟
    <- sem  // up(V原语)
    
    count++
    if count == 10 {
        lock <- true
    }
}

func main() {
    ch := make(chan int, 2) // 带2个int缓冲的Channel(异步Channel，在读完或写满之前都不会阻塞)
    lock := make(chan bool)
    for i := 0; i < 10; i++ {
        go Worker(ch, lock, i)
    }
    <- lock
}