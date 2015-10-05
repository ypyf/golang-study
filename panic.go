package main

import (
	"fmt"
)

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	// 必须放在可能发生panic的代码之前，否则在panic发生时并不会执行到defer中的代码
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g()
	fmt.Println("Returned normally from g.")

}

func g() {
	panic("ERROR")
}
