package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    } else {
        return math.Sqrt(x), nil
    }
}

func find(ch string) (string, error) {
    if ch == "y" {
        return ch, nil
    } else {
        return "", fmt.Errorf("错误")
    }
}

func main() {
    r, e := Sqrt(2)
    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println(r)
    }
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
    ch, _ := find("y")
    fmt.Println(ch)
}
