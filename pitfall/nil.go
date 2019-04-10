package main

import "fmt"

type mytype struct{}

func (a *mytype) Error() string {
	return "mytype error"
}

// 空指针转换为接口（error是一个接口类型）后，并不是nil
// 因为接口类型底层包含两个字段：对象的值和对象的类型
// 只有当这两个字段均为nil时，接口才是nil
// 显然a的类型并不是nil而是mytype
// 总结：当需要返回nil时，直接使用nil字面量或者赋值为nil的接口，不能使用赋值为nil的非接口类型
func foo() error {
	var a *mytype = nil
	return a
}

func main() {
	if foo() != nil {
		fmt.Println("居然不为空!")
	}
}
