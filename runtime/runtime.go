package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Version: %v\n", runtime.Version())
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("GOROOT: %s\n", runtime.GOROOT())
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
}
