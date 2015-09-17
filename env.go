package main

import "fmt"
import "os"
import "strings"

func main() {
	for _, e := range os.Environ() {
		s := strings.Split(e, "=")
		fmt.Printf("%s = %s\n", s[0], s[1])
	}
}
