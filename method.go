package main

import "fmt"

type MyInt int

func (i MyInt)Add(a int) MyInt {
    return MyInt(int(i) + a)
}

func main() {
    a := MyInt(1)
    b := a.Add(2)
    fmt.Printf("1 + 2 = %v\n", b)
}