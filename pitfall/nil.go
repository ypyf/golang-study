package main

import "fmt"

type mytype struct{}

func (a *mytype) Error() string {
	return "mytype error"
}

func foo() error {
	var a *mytype = nil
	return a
}

func main() {

	if foo() != nil {
		fmt.Println("居然不为空!")
	}
}
