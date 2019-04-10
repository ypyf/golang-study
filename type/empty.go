package main

import "fmt"

type empty struct{}

func main() {
	a := empty{}
	b := struct{}{}
	b = a // 字段相同的结构类型是等价的
	fmt.Printf("a = %+v, b = %+v\n", a, b)
}
