package main

import (
	"fmt"
	"runtime"
	"time"
)

//go:noinline
func callAdd() {
	add(3, 5)
}

//go:noinline
func add(a, b int) int {
	return a + b
}

func deadloop() {
	for {
		// 非叶子节点的函数调用有机会调度
		// 因为非叶子节点的调用可能会扩大栈空间，因此编译器会插入morestack()
		callAdd()
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}
 