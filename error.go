package main

import (
	"errors"
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

func returnError() error {
	return FooError
}

var (
	FooError = errors.New("Foo error")
)

func main() {
	r, e := Sqrt(2)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(r)
	}
	err := returnError()
	if err != nil {
		if err == FooError {
			println("foo error")
		}
	}
}
