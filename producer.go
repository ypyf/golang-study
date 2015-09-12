package main

import "fmt"
import "time"

func main() {
  ch := make(chan int)
  done := make(chan struct{})

  go func() {
    for i := 0; i < 10; i++ {
      ch <- i
    }
    done <- struct{}{}
  }()

  go func() {
    for {
      fmt.Println(<-ch)
      time.Sleep(1*time.Second)
    }
  }()

  <-done
}
