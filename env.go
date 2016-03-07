package main

import "fmt"
import "os"
import "strings"

func main() {
	// 获取单个变量
	foo := os.Getenv("FOO")
	fmt.Println("FOO:", foo)

	// 获取全部变量
	for _, e := range os.Environ() {
		s := strings.Split(e, "=")
		fmt.Printf("%s = %s\n", s[0], s[1])
	}
}
