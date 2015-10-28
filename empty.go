package main

import "fmt"

type empty struct{}

func main() {
	a := empty{}
	b := struct{}{}
	fmt.Printf("a = %+v, b = %+v\n", a, b)
}
