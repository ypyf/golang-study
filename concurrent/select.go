package main

import (
	"os"
	"runtime"
	"syscall"
	"time"
)

// s <- "sss"被永远阻塞
func foo1(ch <-chan string) string {
	s := make(chan string)
	select {
	case s <- "sss":
		return <-s
	case v := <-ch:
		return v
	}
}

// 两个case随机执行
func foo2(ch <-chan string) string {
	s := make(chan string, 1)
	select {
	case s <- "sss":
		return <-s
	case v := <-ch:
		return v
	}
}

func main() {
	a := make(chan string, 1)
	a <- "vvv"
	println(foo2(a))
	println(os.Environ()[0])
	runtime.LockOSThread()
	pid := syscall.Getpid()
	syscall.Kill(pid, syscall.SIGTERM)
	// select{}会阻塞goroutine，但不会像for{}那样阻止调度，而是让出CPU资源
	// for{}虽然会阻止调度，但不会阻塞当前goroutine
	go func() {
		//for {
		println("running")
		time.Sleep(2 * time.Second)
		//}
	}()
	for {
	}
}
