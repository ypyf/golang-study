package main

import "fmt"
import "regexp"

const text = `abc
efg
123`

func main() {
	reg := regexp.MustCompile("(?m)^\\w+")
	result := reg.FindAllString(text, -1)
	for _, w := range result {
		fmt.Println(w)
	}
}
