package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

// select 会阻塞goroutine，但不会阻止调度，而是会让出P
func sched() {
	select {
	}
}

// for 会阻止调度，但不会阻塞当前goroutine
// 除非主动调用 runtime.Gosched()
func loop() {
	for {
	}
}

// 这个goroutine将阻塞在系统调用
// P将会与这个阻塞的M分离，重新与另一个M结合
func block() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func main() {
	ch := make(chan int)
	// 设置P的数量
	// P代表了最多可以同时被调度的goroutine的数量（包括main goroutine）
	// 也就是能够同时执行用户级Go代码的操作系统线程的数量
	// 这并不限制能够阻塞在系统调用上的线程数量
	runtime.GOMAXPROCS(2)
	go sched()
	go block()
	go loop()
	<-ch
}
