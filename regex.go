package main

import (
	"fmt"
	"regexp"
)

const text = `abc
efg
123`

func main() {
	reg := regexp.MustCompile(`(?m)^\w+`)
	for _, w := range reg.FindAllString(text, -1) {
		fmt.Println(w)
	}
}
