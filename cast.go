package main

import "fmt"

func main() {
	var a interface{}
	a = nil
	t, _ := a.(string)
	// 当转换失败或a==nil时，下面的语句将会导致panic
	// t := a.(string)
	println(t == "")

	var b interface{}
	b = 123
	p, _ := b.(string)
	println(p == "")

	var m map[string]string = nil
	if _, ok := m["ss"]; !ok {
		println("empty")
	}

	type F struct {
		A int
		B string
	}

	var X struct {
		A int
		C rune
	}
	var xx interface{} = &X
	ff, _ := xx.(*F)
	if ff == nil {
		fmt.Printf("type conversion failed: expected is *component.AuthorizationMessage, got %T\n", xx)
	} else {
		fmt.Printf("%v\n", ff)
	}
}
