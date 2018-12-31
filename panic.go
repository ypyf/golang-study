package main

import (
	"fmt"
)

func main() {
	shouldRecover()
	fmt.Println("Returned normally from shouldRecover().")
}

func shouldRecover() {
	// 必须放在可能发生panic的代码之前，否则在panic发生时并不会执行到defer中的代码
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in shouldRecover(): %s\n", r)
		}
	}()
	fmt.Println("Calling g.")
	mayPanic()
	fmt.Println("Returned normally from mayPanic().")

}

func mayPanic() {
	panic("ERROR")
}
