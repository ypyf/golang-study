package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("当前版本：%v\n", runtime.Version())
}
