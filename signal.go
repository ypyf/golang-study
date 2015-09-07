package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
    go func(x ...int) {
        fmt.Printf("%v\n", x)
    }(1,2,3)
    
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT)
    <-ch
    fmt.Println("Program exit.")
}