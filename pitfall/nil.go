package main

import "fmt"

type mytype struct{}

func (a *mytype) Error() string {
	return "mytype error"
}

// 空指针转换为接口（error）后，并不是nil
// 因为接口类型底层包含两个字段：对象的值和对象的类型
// 只有当这两个字段均为nil时，接口才是nil
// 显然下面代码中的a的类型并不是nil而是mytype
// 总结：当需要返回nil时，直接使用nil字面量
func foo() error {
	var a *mytype = nil
	return a
}

func main() {

	if foo() != nil {
		fmt.Println("居然不为空!")
	}
}
