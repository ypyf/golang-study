package main

import (
	"fmt"
	"unsafe"
)

type Message struct{}

type FooMessage struct {
	Message
	X string
}

func foo(msg *Message) {

}

func main() {
	var s struct{}
	fmt.Println(unsafe.Sizeof(s))

	var a, b struct{}
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&a == &b)
	msg := FooMessage{X: "hello"}
	foo(&msg.Message)
}
