package main

import (
	"fmt"
	"strconv"
)

// int64转string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// int32转string
func Int32ToString(i int32) string {
	return Int64ToString(int64(i))
}

func MathFloor(a int32, b int32) float64 {
	if b == 0 {
		return 0
	}
	c := Int32ToString(a)
	d := Int32ToString(b)
	e, _ := strconv.ParseFloat(c, 32/64)
	f, _ := strconv.ParseFloat(d, 32/64)
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", e/f*100), 64)
	if value < 0 {
		return 0
	}
	return value
}

func MathFloor2(a int32, b int32) float64 {
	return float64(a) / float64(b)
}

func main() {
	fmt.Printf("11/24=%f\n", MathFloor(11, 24))
	fmt.Printf("11/24=%.4f\n", MathFloor2(11, 24))
}
