package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))

	var a, b struct{}
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&a == &b)
}
