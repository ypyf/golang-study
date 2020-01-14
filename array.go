package main

import "fmt"

func foo(arr [4]int) {
	fmt.Printf("%v\n", arr)
}

func main() {
	a := [4]int{1, 2, 3, 4}
	foo(a)
}
