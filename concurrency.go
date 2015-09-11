package main

import "fmt"
import "time"

func main() {
    v := []rune("Go Programming Language")
    
    go func() {
        v[0] = 'X'
        time.Sleep(1*time.Second)
        v[1] = 'Z'
    }()
    
    go func() {
        fmt.Println(string(v))  // v = ?
    }()
    
    v[0] = 'S'
    time.Sleep(1*time.Second)
    v[1] = 'B'
    
    
    fmt.Println(string(v))  // v = ?
}