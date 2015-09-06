package main

import "fmt"

// RandomBits函数 返回一个管道，用于产生一个比特随机序列
func randomBits() <-chan int {
    ch := make(chan int)
    go func() {
        for {
            // 当两个send语句都可以处理时，select会随机选择一个
            select {
            case ch <- 0: 
            case ch <- 1:
            }
        }
    }()
    return ch
}

func main() {
    for i := 0; i < 100; i++ {
        fmt.Printf("%d", <- randomBits())
    }
    fmt.Println()
}
