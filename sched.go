package main

import (
	"runtime"
	"time"
)

func sched() {
	for {
		runtime.Gosched()
	}
}

// 运行后该goroutine将无法被调离
// 除非主动调用 runtime.Gosched()
func loop() {
	for {}
}

func main() {
	// 设置P的数量
	// P代表了最多可以同时被调度的goroutine的数量（包括main goroutine）
	runtime.GOMAXPROCS(2)
	go sched()
	go loop()
	time.Sleep(time.Second)
}
