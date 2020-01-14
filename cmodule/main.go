package main

import "C"
import "fmt"

//export Sum
func Sum(a, b int) int {
	return a + b
}

func subset(a []int) (r [][]int) {
	if len(a) == 0 {
		r = append(r, []int{})
		return
	}
	for _, s := range subset(a[1:]) {
		r = append(r, s)
		// 跳过空集
		if len(s) > 0 {
			s = append(s, a[0])
			r = append(r, s)
		}
	}
	r = append(r, []int{a[0]})
	return
}

func main() {
	for _, x := range subset([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		fmt.Printf("%v\n", x)
	}
}
