package main

import (
	"fmt"
	"strings"
)

func removeElement() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = append(a[:2], a[3:]...)
	fmt.Printf("%v\n", a)
}

func main() {
	s0 := make([]int32, 10)
	s1 := append(s0, 12)
	fmt.Printf("len %d cap %d\n%v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 12)
	fmt.Printf("len %d cap %d\n%v\n", len(s1), cap(s1), s1)
	ps0 := &s0
	pps0 := &ps0
	fmt.Printf("ps0 %v\npps0 %v\n", *ps0, **pps0)

	// array
	arr := [...]string{}
	fmt.Printf("type:%T, arr.len = %d, arr.cap = %d\n", arr, len(arr), cap(arr))

	arr2 := [...]string{"hello", "world"}
	fmt.Printf("arr2.len = %d, arr2.cap = %d\n", len(arr2), cap(arr2))

	// 数组转换为slice
	slice := new([2]int32)[:]
	slice[0], slice[1] = 123, 456
	fmt.Printf("slice = %v\n", slice)

	// nested slice
	a, b := []int{1, 2}, []int{3, 4}
	c := [][]int{a, b}
	fmt.Println(c)

	removeElement()

	// Join slice
	scopes := []string{"email", "sms"}
	fmt.Println(strings.Join(scopes, ","))
}
